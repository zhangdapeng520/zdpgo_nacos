/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package naming_client

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/zhangdapeng520/zdpgo_nacos/nacos/clients/cache"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/model"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/util"
)

type HostReactor struct {
	serviceInfoMap       cache.ConcurrentMap
	cacheDir             string
	updateThreadNum      int
	serviceProxy         NamingProxy
	pushReceiver         PushReceiver
	subCallback          SubscribeCallback
	updateTimeMap        cache.ConcurrentMap
	updateCacheWhenEmpty bool
}

const Default_Update_Thread_Num = 20

func NewHostReactor(serviceProxy NamingProxy, cacheDir string, updateThreadNum int, notLoadCacheAtStart bool, subCallback SubscribeCallback, updateCacheWhenEmpty bool) HostReactor {
	if updateThreadNum <= 0 {
		updateThreadNum = Default_Update_Thread_Num
	}
	hr := HostReactor{
		serviceProxy:         serviceProxy,
		cacheDir:             cacheDir,
		updateThreadNum:      updateThreadNum,
		serviceInfoMap:       cache.NewConcurrentMap(),
		subCallback:          subCallback,
		updateTimeMap:        cache.NewConcurrentMap(),
		updateCacheWhenEmpty: updateCacheWhenEmpty,
	}
	pr := NewPushReceiver(&hr)
	hr.pushReceiver = *pr
	if !notLoadCacheAtStart {
		hr.loadCacheFromDisk()
	}
	go hr.asyncUpdateService()
	return hr
}

func (hr *HostReactor) loadCacheFromDisk() {
	serviceMap := cache.ReadServicesFromFile(hr.cacheDir)
	if serviceMap == nil || len(serviceMap) == 0 {
		return
	}
	for k, v := range serviceMap {
		hr.serviceInfoMap.Set(k, v)
	}
}

func (hr *HostReactor) ProcessServiceJson(result string) {
	service := util.JsonToService(result)
	if service == nil {
		return
	}
	cacheKey := util.GetServiceCacheKey(service.Name, service.Clusters)

	oldDomain, ok := hr.serviceInfoMap.Get(cacheKey)
	if ok && !hr.updateCacheWhenEmpty {
		//if instance list is empty,not to update cache
		if service.Hosts == nil || len(service.Hosts) == 0 {
			return
		}
	}
	hr.updateTimeMap.Set(cacheKey, uint64(util.CurrentMillis()))
	hr.serviceInfoMap.Set(cacheKey, *service)
	if !ok || ok && !reflect.DeepEqual(service.Hosts, oldDomain.(model.Service).Hosts) {
		cache.WriteServicesToFile(*service, hr.cacheDir)
		hr.subCallback.ServiceChanged(service)
	}
}

func (hr *HostReactor) GetServiceInfo(serviceName string, clusters string) (model.Service, error) {
	key := util.GetServiceCacheKey(serviceName, clusters)
	cacheService, ok := hr.serviceInfoMap.Get(key)
	if !ok {
		hr.updateServiceNow(serviceName, clusters)
		if cacheService, ok = hr.serviceInfoMap.Get(key); !ok {
			return model.Service{}, errors.New("get service info failed")
		}
	}

	return cacheService.(model.Service), nil
}

func (hr *HostReactor) GetAllServiceInfo(nameSpace, groupName string, pageNo, pageSize uint32) model.ServiceList {
	data := model.ServiceList{}
	result, err := hr.serviceProxy.GetAllServiceInfoList(nameSpace, groupName, pageNo, pageSize)
	if err != nil {
		return data
	}
	if result == "" {
		return data
	}

	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		return data
	}
	return data
}

func (hr *HostReactor) updateServiceNow(serviceName, clusters string) {
	result, err := hr.serviceProxy.QueryList(serviceName, clusters, hr.pushReceiver.port, false)

	if err != nil {
		return
	}
	if result == "" {
		return
	}
	hr.ProcessServiceJson(result)
}

func (hr *HostReactor) asyncUpdateService() {
	sema := util.NewSemaphore(hr.updateThreadNum)
	for {
		for _, v := range hr.serviceInfoMap.Items() {
			service := v.(model.Service)
			lastRefTime, ok := hr.updateTimeMap.Get(util.GetServiceCacheKey(service.Name, service.Clusters))
			if !ok {
				lastRefTime = uint64(0)
			}
			if uint64(util.CurrentMillis())-lastRefTime.(uint64) > service.CacheMillis {
				sema.Acquire()
				go func() {
					hr.updateServiceNow(service.Name, service.Clusters)
					sema.Release()
				}()
			}
		}
		time.Sleep(1 * time.Second)
	}
}
