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
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zhangdapeng520/zdpgo_nacos/nacos/clients/cache"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/common/constant"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/model"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/util"
	nsema "github.com/zhangdapeng520/zdpgo_nacos/toolkits_concurrent/semaphore"
)

type BeatReactor struct {
	beatMap             cache.ConcurrentMap
	serviceProxy        NamingProxy
	clientBeatInterval  int64
	beatThreadCount     int
	beatThreadSemaphore *nsema.Semaphore
	beatRecordMap       cache.ConcurrentMap
	mux                 *sync.Mutex
}

const Default_Beat_Thread_Num = 20

func NewBeatReactor(serviceProxy NamingProxy, clientBeatInterval int64) BeatReactor {
	br := BeatReactor{}
	if clientBeatInterval <= 0 {
		clientBeatInterval = 5 * 1000
	}
	br.beatMap = cache.NewConcurrentMap()
	br.serviceProxy = serviceProxy
	br.clientBeatInterval = clientBeatInterval
	br.beatThreadCount = Default_Beat_Thread_Num
	br.beatRecordMap = cache.NewConcurrentMap()
	br.beatThreadSemaphore = nsema.NewSemaphore(br.beatThreadCount)
	br.mux = new(sync.Mutex)
	return br
}

func buildKey(serviceName string, ip string, port uint64) string {
	return serviceName + constant.NAMING_INSTANCE_ID_SPLITTER + ip + constant.NAMING_INSTANCE_ID_SPLITTER + strconv.Itoa(int(port))
}

func (br *BeatReactor) AddBeatInfo(serviceName string, beatInfo model.BeatInfo) {

	k := buildKey(serviceName, beatInfo.Ip, beatInfo.Port)
	defer br.mux.Unlock()
	br.mux.Lock()
	if data, ok := br.beatMap.Get(k); ok {
		beatInfo := data.(*model.BeatInfo)
		atomic.StoreInt32(&beatInfo.State, int32(model.StateShutdown))
		br.beatMap.Remove(k)
	}
	br.beatMap.Set(k, &beatInfo)
	go br.sendInstanceBeat(k, &beatInfo)
}

func (br *BeatReactor) RemoveBeatInfo(serviceName string, ip string, port uint64) {

	k := buildKey(serviceName, ip, port)
	defer br.mux.Unlock()
	br.mux.Lock()
	data, exist := br.beatMap.Get(k)
	if exist {
		beatInfo := data.(*model.BeatInfo)
		atomic.StoreInt32(&beatInfo.State, int32(model.StateShutdown))
	}
	br.beatMap.Remove(k)
}

func (br *BeatReactor) sendInstanceBeat(k string, beatInfo *model.BeatInfo) {
	for {
		br.beatThreadSemaphore.Acquire()
		//如果当前实例注销，则进行停止心跳
		if atomic.LoadInt32(&beatInfo.State) == int32(model.StateShutdown) {

			br.beatThreadSemaphore.Release()
			return
		}

		//进行心跳通信
		beatInterval, err := br.serviceProxy.SendBeat(*beatInfo)
		if err != nil {

			br.beatThreadSemaphore.Release()
			t := time.NewTimer(beatInfo.Period)
			<-t.C
			continue
		}
		if beatInterval > 0 {
			beatInfo.Period = time.Duration(time.Millisecond.Nanoseconds() * beatInterval)
		}

		br.beatRecordMap.Set(k, util.CurrentMillis())
		br.beatThreadSemaphore.Release()

		t := time.NewTimer(beatInfo.Period)
		<-t.C
	}
}
