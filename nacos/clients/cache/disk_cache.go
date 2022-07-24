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

package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/go-errors/errors"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/common/constant"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/common/file"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/model"
	"github.com/zhangdapeng520/zdpgo_nacos/nacos/util"
)

func GetFileName(cacheKey string, cacheDir string) string {

	if runtime.GOOS == constant.OS_WINDOWS {
		cacheKey = strings.ReplaceAll(cacheKey, ":", constant.WINDOWS_LEGAL_NAME_SPLITER)
	}

	return cacheDir + string(os.PathSeparator) + cacheKey
}

func WriteServicesToFile(service model.Service, cacheDir string) {
	file.MkdirIfNecessary(cacheDir)
	sb, _ := json.Marshal(service)
	domFileName := GetFileName(util.GetServiceCacheKey(service.Name, service.Clusters), cacheDir)

	err := ioutil.WriteFile(domFileName, sb, 0666)
	if err != nil {

	}

}

func ReadServicesFromFile(cacheDir string) map[string]model.Service {
	files, err := ioutil.ReadDir(cacheDir)
	if err != nil {

		return nil
	}
	serviceMap := map[string]model.Service{}
	for _, f := range files {
		fileName := GetFileName(f.Name(), cacheDir)
		b, err := ioutil.ReadFile(fileName)
		if err != nil {

			continue
		}

		s := string(b)
		service := util.JsonToService(s)

		if service == nil {
			continue
		}

		serviceMap[f.Name()] = *service
	}

	return serviceMap
}

func WriteConfigToFile(cacheKey string, cacheDir string, content string) {
	file.MkdirIfNecessary(cacheDir)
	fileName := GetFileName(cacheKey, cacheDir)
	err := ioutil.WriteFile(fileName, []byte(content), 0666)
	if err != nil {

	}
}

func ReadConfigFromFile(cacheKey string, cacheDir string) (string, error) {
	fileName := GetFileName(cacheKey, cacheDir)
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to read config cache file:%s,err:%+v ", fileName, err))
	}
	return string(b), nil
}
