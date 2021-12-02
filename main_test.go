package zgo_nacos

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"zgo_config"
	"zgo_nacos/test/config"
)

func TestConfig(t *testing.T) {
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "1a67e97c-fa54-44dd-961b-4dcaeed9217c", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
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

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user_web.json",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml
	serverConfig := config.ServerConfig{}
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	time.Sleep(5 * time.Second)

}

// 使用方法获取配置
func TestConfig1(t *testing.T) {
	client := GetNacosClient("127.0.0.1", 8848, "1a67e97c-fa54-44dd-961b-4dcaeed9217c")
	content := GetNacosContent(client, "user_web.json", "dev")
	fmt.Println(content)
	serverConfig := config.ServerConfig{}
	ParseNacosJsonConfig(&serverConfig, content)
	fmt.Println(serverConfig)
}

// 测试获取用户服务配置
func TestGetServiceUserConfig(t *testing.T) {
	client := GetNacosClient("127.0.0.1", 8848, "1a67e97c-fa54-44dd-961b-4dcaeed9217c")
	content := GetNacosContent(client, "user_service.json", "dev")
	fmt.Println(content)
	serverConfig := zgo_config.ServiceConfig{}
	ParseNacosJsonConfig(&serverConfig, content)
	fmt.Println(serverConfig.Mysql)
}
