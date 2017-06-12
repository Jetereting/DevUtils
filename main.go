package main

import (
	_ "beegoAutoDoc/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	runMode:=beego.AppConfig.String("runmode")
	host := beego.AppConfig.String(runMode+"::host")
	user := beego.AppConfig.String(runMode+"::user")
	database := beego.AppConfig.String(runMode+"::database")
	pwd := beego.AppConfig.String(runMode+"::pwd")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", user + ":" + pwd + "@tcp(" + host + ")/" + database)
}

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}

