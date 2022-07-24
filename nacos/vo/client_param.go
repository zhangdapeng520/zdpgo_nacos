package vo

import "github.com/zhangdapeng520/zdpgo_nacos/nacos/common/constant"

type NacosClientParam struct {
	ClientConfig  *constant.ClientConfig  // optional
	ServerConfigs []constant.ServerConfig // optional
}
