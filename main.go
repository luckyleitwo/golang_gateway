package main

import (
	"golang_gateway/router"
	"golang_gateway/utils"
)

func main() {
	// 初始化 config 文件
	utils.InitConfig()
	// 初始化 MySql
	utils.InitMySql()
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
