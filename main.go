package main

import (
	_ "myblog/routers"

	"myblog/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}
func main() {
	beego.Run()
}
