package main

import (
	"fmt"

	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./cfg/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	cfg.GlobalConfig = &cfg.Config{}
	if err := viper.Unmarshal(cfg.GlobalConfig); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("load config succ: \n%+v\n", cfg.GlobalConfig)
}

func main() {

}
