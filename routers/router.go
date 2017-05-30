// @APIVersion 1.0.0
// @Title 便于开发的工具集合
// @Description nothing to talk.
// @Contact alloyjetereting@gmail.com
package routers

import (
	"beegoAutoDoc/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/excelToMysql",
			beego.NSInclude(&controllers.ExcelToMysqlController{})))
	beego.AddNamespace(ns)
}
