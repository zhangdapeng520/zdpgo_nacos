package zdpgo_nacos

// MysqlConfig MySQL配置信息
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`         // 主机地址
	Port     int    `mapstructure:"port" json:"port"`         // 端口号
	Name     string `mapstructure:"db" json:"database"`       // 数据库名称
	User     string `mapstructure:"user" json:"username"`     // 用户名
	Password string `mapstructure:"password" json:"password"` // 密码
}

// ConsulConfig consul配置
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"` // 主机
	Port int    `mapstructure:"port" json:"port"` // 端口号
}

// ServiceConfig service配置
type ServiceConfig struct {
	Name       string       `mapstructure:"name" json:"name"`     // 服务名称
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`   // MySQL配置
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"` // consul配置
}

// Config nacos配置
type Config struct {
	Host      string `mapstructure:"host"`      // 主机地址
	Port      uint64 `mapstructure:"port"`      // 主机端口号
	Namespace string `mapstructure:"namespace"` // 名称空间
	User      string `mapstructure:"user"`      // 用户
	Password  string `mapstructure:"password"`  //密码
	DataId    string `mapstructure:"dataid"`    // dataid
	Group     string `mapstructure:"group"`     // 组
}
