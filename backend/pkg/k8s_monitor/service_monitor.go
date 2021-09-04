// Copyright © 2021 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s_monitor

import (
	"sync"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type ServiceMonitor struct {
	serviceIpMap *sync.Map // Hold cluster IPs of services
	clientset    kubernetes.Interface
	stopCh       chan struct{}
}

func CreateServiceMonitor(clientset kubernetes.Interface) (*ServiceMonitor, error) {
	stopCh := make(chan struct{})
	var serviceIpMap sync.Map
	return &ServiceMonitor{
		serviceIpMap: &serviceIpMap,
		clientset:    clientset,
		stopCh:       stopCh,
	}, nil
}

func (m *ServiceMonitor) Start() {
	log.Info("Starting Service monitor")
	watchlist := cache.NewListWatchFromClient(m.clientset.CoreV1().RESTClient(), v1.ResourceServices.String(), v1.NamespaceAll, fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Service{},
		ResyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    m.addService,
			DeleteFunc: m.deleteService,
			UpdateFunc: m.updateService,
		},
	)
	go controller.Run(m.stopCh)
}

func (m *ServiceMonitor) Stop() {
	log.Info("Stopping Service monitor")
	close(m.stopCh)
}

func (m *ServiceMonitor) addService(obj interface{}) {
	service := obj.(*v1.Service)
	m.serviceIpMap.Store(service.Spec.ClusterIP, true)

	log.Tracef("Service added: service=%+v (%v)", service.Name+"."+service.Namespace, service.Spec.ClusterIP)
}

func (m *ServiceMonitor) deleteService(obj interface{}) {
	service := obj.(*v1.Service)
	m.serviceIpMap.Delete(service.Spec.ClusterIP)
	log.Tracef("Service deleted: service=%+v (%v)", service.Name+"."+service.Namespace, service.Spec.ClusterIP)
}

func (m *ServiceMonitor) updateService(oldObj, newObj interface{}) {
	oldService := oldObj.(*v1.Service)
	newService := newObj.(*v1.Service)

	if oldService.Spec.ClusterIP != newService.Spec.ClusterIP {
		m.serviceIpMap.Delete(oldService.Spec.ClusterIP)
		m.serviceIpMap.Store(newService.Spec.ClusterIP, true)
	}

	log.Tracef("Service updated: old service=service=%+v (%v), new service=service=%+v (%v)",
		oldService.Name+"."+oldService.Namespace, oldService.Spec.ClusterIP,
		newService.Name+"."+newService.Namespace, newService.Spec.ClusterIP,
	)
}

func (m *ServiceMonitor) IsServiceIP(ipStr string) bool {
	ret := false
	m.serviceIpMap.Range(func(key, value interface{}) bool {
		if key == ipStr {
			ret = true
			// stop iteration
			return false
		}
		// continue iteration
		return true
	})

	return ret
}
