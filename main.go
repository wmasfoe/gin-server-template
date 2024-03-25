package main

import (
	"fmt"
	"go-blog-server/core"
	"go-blog-server/global"
)

func main() {
	core.InitConf()
	core.InitGorm()
	core.InitLogger()

	fmt.Println(global.CONFIG)
	fmt.Println(global.DB)
	global.Log.Info("异常")
}
