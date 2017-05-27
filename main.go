package main

import (
	_ "beegoAutoDoc/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	host :=beego.AppConfig.String("mysql::host")
	user :=beego.AppConfig.String("mysql::user")
	database :=beego.AppConfig.String("mysql::database")
	pwd :=beego.AppConfig.String("mysql::pwd")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", user+":"+pwd+"@tcp("+host+")/"+database)
	//orm.RegisterDataBase("default", "mysql","qdxg:qdxg_603@tcp(192.168.1.201:3306)/qdxg")
}


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

