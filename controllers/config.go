package controllers

import (
	"github.com/astaxie/beego"
)




// 配置文件
type ConfigController struct {
	beego.Controller
}

// @Summary Mysql 配置
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /mysql [get]
func (c *ConfigController) MysqlDetail(){
	a:=`……
	`+
	`
	host:`+beego.AppConfig.String("host")+
	`
	user:`+beego.AppConfig.String("user")+
	`
	database:`+beego.AppConfig.String("database")
	c.Ctx.WriteString(a)
}