package zdpgo_nacos

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zhangdapeng520/zdpgo_log"
)

// Nacos nacos核心对象
type Nacos struct {
	log    *zdpgo_log.Log              // 日志对象
	config *NacosConfig                // 配置对象
	client config_client.IConfigClient // nacos配置客户端对象
}

type NacosConfig struct {
	Debug       bool   // 是否为debug模式
	LogFilePath string // 日志文件路径
	Host        string // nacos主机地址
	Port        uint16 // nacos端口号
	NamespaceID string // 名称空间id
}

func New(config NacosConfig) *Nacos {
	n := Nacos{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_nacos.log"
	}
	l := zdpgo_log.New(zdpgo_log.LogConfig{
		Debug: config.Debug,
		Path:  config.LogFilePath,
	})
	n.log = l

	// 校验参数
	if config.Host == "" {
		n.log.Warning("nacos主机地址Host为空，将使用默认值：127.0.0.1")
		n.config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		n.log.Warning("nacos端口号Port为空，将使用默认值：8848")
		config.Port = 8848
	}
	if config.NamespaceID == "" {
		n.log.Panic("名称空间不能为空！")
	}

	// 配置
	n.config = &config

	// 初始化客户端
	n.initClient()

	return &n
}

// InitClient 初始化nacos客户端
func (n *Nacos) initClient() {

	// 约束配置
	sc := []constant.ServerConfig{
		{
			IpAddr: n.config.Host,
			Port:   uint64(n.config.Port),
		},
	}

	// 客户端配置
	cc := constant.ClientConfig{
		NamespaceId:         n.config.NamespaceID,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// 配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		n.log.Error("获取nacos配置失败：", err.Error())
	}

	// 初始化客户端
	n.client = configClient
}

// GetContent 获取nacos配置信息
// @param dataId 配置组id
// @param group 配置组名称
// @return content 配置内容
func (n *Nacos) GetContent(dataId, group string) (content string) {
	n.log.Info("nacos配置组：", dataId, group)
	content, err := n.client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if err != nil {
		n.log.Error("读取指定配置失败：", err)
	}
	return content
}

// ParseJsonConfig 解析Json配置
// @param config 配置对象
// @param jsonContent json配置内容
func (n *Nacos) ParseJsonConfig(config interface{}, jsonContent string) {
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err := json.Unmarshal([]byte(jsonContent), &config)
	if err != nil {
		n.log.Error("解析json配置失败：", err)
	}
}

// ListenConfig 监听nacos配置
// @param group 配置组名称
// @param content 配置内容
// @param config 配置对象
func (n *Nacos) ListenConfig(dataId, group string, config interface{}) error {
	err := n.client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			content := n.GetContent(dataId, group)
			n.ParseJsonConfig(config, content)
		},
	})
	return err
}
