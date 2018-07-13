package main

import (
	_ "DevUtils/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"github.com/astaxie/beego/orm"
)

func init() {
	host := beego.AppConfig.String("host")
	user := beego.AppConfig.String("user")
	database := beego.AppConfig.String("database")
	pwd := beego.AppConfig.String("pwd")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", user + ":" + pwd + "@tcp(" + host + ")/" + database)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}

