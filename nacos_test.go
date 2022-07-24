package zdpgo_nacos

import (
	"fmt"
	"testing"
)

func prepareNacos() *Nacos {
	n := New(NacosConfig{
		Debug:       true,
		Host:        "127.0.0.1",
		Port:        8848,
		NamespaceID: "283df317-fa59-46db-9cac-838970599ef4",
	})
	return n
}

// 测试新建
func TestNacos_New(t *testing.T) {
	n := prepareNacos()
	fmt.Println(n)
}

// 测试初始化客户端
func TestNacos_InitClient(t *testing.T) {
	n := prepareNacos()
	fmt.Println(n)
	fmt.Println(n.client)
}

// 测试获取配置内容
func TestNacos_GetContent(t *testing.T) {
	n := prepareNacos()
	content := n.GetContent("user_web.json", "dev")
	fmt.Println(content)
	fmt.Println(string(content))
}

// 测试解析json配置
func TestNacos_ParseJsonConfig(t *testing.T) {
	n := prepareNacos()
	content := n.GetContent("user_web.json", "dev")

	config := ServiceConfig{}
	n.ParseJsonConfig(&config, content)
	fmt.Println(config)
	fmt.Println(config.Name)
	fmt.Println(config.ConsulInfo)
	fmt.Println(config.MysqlInfo)
}
