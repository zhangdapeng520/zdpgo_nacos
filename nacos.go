package zgo_nacos

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// 获取nacos客户端
func GetNacosClient(nacosIp string, nacosPort uint16, namespaceId string) config_client.IConfigClient {
	sc := []constant.ServerConfig{
		{
			IpAddr: nacosIp,
			Port:   uint64(nacosPort),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	return configClient
}

// 获取nacos配置信息
func GetNacosContent(client config_client.IConfigClient, dataId, group string) string {
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if err != nil {
		panic(err)
	}
	return content
}

// 解析Json配置
func ParseNacosJsonConfig(config interface{}, jsonContent string) {
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err := json.Unmarshal([]byte(jsonContent), &config)
	if err != nil {
		panic(err)
	}
}

// 监听nacos配置
func ListenNacosConfig(client config_client.IConfigClient, dataId, group string, config interface{}) error {
	err := client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			content := GetNacosContent(client, dataId, group)
			ParseNacosJsonConfig(config, content)
		},
	})
	return err
}
