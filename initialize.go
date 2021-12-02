package zgo_nacos

import (
	"zgo_config"
)

var (
	// nacos配置信息
	NacosConfig = &zgo_config.NacosConfig{}

	// service配置信息
	ServiceConfig = &zgo_config.ServiceConfig{}

	// web配置信息
	WebConfig = &zgo_config.WebConfig{}
)

// 读取nacos配置信息
func InitNacosConfig() {
	zgo_config.InitDefaultConfig(&NacosConfig)
}

// 从nacos读取service信息
func InitNacosServiceConfig() {
	InitNacosConfig()
	client := GetNacosClient(
		NacosConfig.Host,
		uint16(NacosConfig.Port),
		NacosConfig.NamespaceId,
	)
	content := GetNacosContent(client, NacosConfig.DataId, NacosConfig.Group)
	ParseNacosJsonConfig(ServiceConfig, content)
}

// 从nacos读取web信息
func InitNacosWebConfig() {
	InitNacosConfig()
	client := GetNacosClient(
		NacosConfig.Host,
		uint16(NacosConfig.Port),
		NacosConfig.NamespaceId,
	)
	content := GetNacosContent(client, NacosConfig.DataId, NacosConfig.Group)
	ParseNacosJsonConfig(WebConfig, content)
}
