package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
	"myblog/utils"
	_ "myblog/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
