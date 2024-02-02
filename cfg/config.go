package cfg

type Config struct {
	Server             ServerConfig            `yaml:"server"`
	Database           DatabaseConfig          `yaml:"database"`
	ProxyServer        ProxyServerConfig       `yaml:"proxyServer"`
	Spiders            map[string]SpiderConfig `yaml:"spiders"`
	ArticlesRequestUrl string                  `yaml:"articlesRequestUrl"`
	ArticleUrl         []ArticleUrlConfig      `yaml:"articleUrl"`
	Logger             LoggerConfig            `yaml:"logger"`
}

type LoggerConfig struct {
	Level      string `yaml:"level"`
	OnSave     bool   `yaml:"onSave"`
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
	DbName       string `yaml:"dbname"`
	ReservedConn int    `yaml:"reservedConn"`
}

type ProxyServerConfig struct {
	Port        int    `yaml:"port"`
	Host        string `yaml:"host"`
	CustomProxy string `yaml:"customProxy"`
}

type SpiderConfig struct {
	Prefix string `yaml:"prefix"`
	Spec   int    `yaml:"spec"`
}

type ArticleUrlConfig struct {
	Referer string   `yaml:"referer"`
	RealmId []string `yaml:"realmId"`
	Spec    int      `yaml:"spec"`
}

func (a *ArticleUrlConfig) Clone() ArticleUrlConfig {
	realmId := make([]string, len(a.RealmId))
	copy(realmId, a.RealmId)

	ret := ArticleUrlConfig{
		Referer: a.Referer,
		RealmId: realmId,
		Spec:    a.Spec,
	}

	return ret
}
