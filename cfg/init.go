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
	// for product
	if err := loadConfigByPath("./cfg/config.yaml"); err == nil {
		correctionConfig()
		log.Printf("load config for product succ: \n%+v\n", GlobalConfig)
		return
	} else {
		log.Printf("load config for product fail: %+v\n", err)
	}

	// for test
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	if err := loadConfigByPath(dir + "/config.yaml"); err != nil {
		log.Fatalf("load config for test fail: %+v\n", err)
		return
	}

	correctionConfig()
	log.Printf("load config for test succ: \n%+v\n", GlobalConfig)
}

func correctionConfig() {
	articlesListConfig := GlobalConfig.Observers["articles"]
	if articlesListConfig.Spec > len(GlobalConfig.ArticleUrl) {
		articlesListConfig.Spec = len(GlobalConfig.ArticleUrl)
	}
	if articlesListConfig.IdleTime <= 0 {
		articlesListConfig.IdleTime = 10
	}
	GlobalConfig.Observers["articles"] = articlesListConfig

	commentsConfig := GlobalConfig.Observers["comments"]
	if commentsConfig.Spec < articlesListConfig.Spec {
		commentsConfig.Spec = articlesListConfig.Spec
	}
	GlobalConfig.Observers["comments"] = commentsConfig
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
