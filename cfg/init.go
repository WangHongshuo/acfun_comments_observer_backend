package cfg

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var GlobalConfig *Config

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	if err := loadConfigByPath(dir + "\\config.yaml"); err != nil {
		log.Fatalf("load config fail: %+v\n", err)
		return
	}

	// correction
	articlesListConfig := GlobalConfig.Observers["articles"]
	if articlesListConfig.Spec > len(GlobalConfig.ArticleUrl) {
		articlesListConfig.Spec = len(GlobalConfig.ArticleUrl)
		GlobalConfig.Observers["articles"] = articlesListConfig
	}

	commentsConfig := GlobalConfig.Observers["comments"]
	if commentsConfig.Spec < articlesListConfig.Spec {
		commentsConfig.Spec = articlesListConfig.Spec
		GlobalConfig.Observers["comments"] = commentsConfig
	}

	log.Printf("load config succ: \n%+v\n", GlobalConfig)
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
