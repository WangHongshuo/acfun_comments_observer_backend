package cfg

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var GlobalConfig *Config

type Config struct {
	Server             ServerConfig       `yaml:"server"`
	Database           DatabaseConfig     `yaml:"database"`
	ProxyServer        ProxyServerConfig  `yaml:"proxyServer"`
	Spiders            []SpiderConfig     `yaml:"spiders"`
	ArticlesRequestUrl string             `yaml:"articlesRequestUrl"`
	ArticleUrl         []ArticleUrlConfig `yaml:"articleUrl"`
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

type SpiderConfig struct {
	Type   string `yaml:"type"`
	Prefix string `yaml:"prefix"`
	Spec   int    `yaml:"spec"`
}

type ArticleUrlConfig struct {
	Referer string   `yaml:"referer"`
	RealmId []string `yaml:"realmId"`
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	if err := loadConfigByPath(dir + "\\config.yaml"); err != nil {
		fmt.Printf("load config fail: %+v\n", err)
		return
	}
	fmt.Printf("load config succ: \n%+v\n", GlobalConfig)
}

func loadConfigByPath(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("load config fail: %v", err)
	}
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		return fmt.Errorf("load config fail: %v", err)
	}
	return nil
}
