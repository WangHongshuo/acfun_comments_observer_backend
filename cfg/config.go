package cfg

var GlobalConfig *Config

type Config struct {
	Server        ServerConfig         `yaml:"server"`
	Database      DatabaseConfig       `yaml:"database"`
	ProxyServer   ProxyServerConfig    `yaml:"proxyServer"`
	SpiderWorkers []SpiderWorkerConfig `yaml:"spiderWorkers"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	UserName  string `yaml:"username"`
	Password  string `yaml:"password"`
	TableName string `yaml:"tablename"`
}

type ProxyServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type SpiderWorkerConfig struct {
	Type   string `yaml:"type"`
	Prefix string `yaml:"prefix"`
	Spec   int    `yaml:"spec"`
}
