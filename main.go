package main

import (
	"fmt"
	"go-blog-server/core"
	"go-blog-server/global"
)

func main() {
	core.InitConf()
	core.InitGorm()

	fmt.Println(global.CONFIG)
	fmt.Println(global.DB)
}
