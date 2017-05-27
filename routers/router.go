// @APIVersion 1.0.0
// @Title 暂时用作excle导入到临时表
// @Description nothing to talk.
// @Contact alloyjetereting@gmail.com
package routers

import (
	"beegoAutoDoc/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/main",
			beego.NSInclude(&controllers.MainController{})))
	beego.AddNamespace(ns)
}
