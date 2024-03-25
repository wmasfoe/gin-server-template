package core

import (
	"fmt"
	"go-blog-server/config"
	"go-blog-server/global"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// 读取 setting.yaml 的数据
func InitConf() {
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get config error %v", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}
	log.Printf("config yaml load success!")

	// 添加到全局变量
	global.CONFIG = c
}
