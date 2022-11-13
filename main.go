package main

import (
	"golang_gateway/router"
	"golang_gateway/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySql()
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
