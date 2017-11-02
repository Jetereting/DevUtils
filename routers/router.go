// @APIVersion 1.1.2
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
		beego.NSNamespace("/config",
			beego.NSInclude(
				&controllers.ConfigController{},
			),
		),beego.NSNamespace("/excelToMysql",
			beego.NSInclude(
				&controllers.ExcelToMysqlController{},
			),
		),beego.NSNamespace("/mysqlToExcel",
			beego.NSInclude(
				&controllers.MysqlToExcelController{},
			),
		),
		beego.NSNamespace("/modelTranslateMarkDown",
			beego.NSInclude(
				&controllers.ModelTranslateMarkDownController{},
			),
		),
		beego.NSNamespace("/toLowercase",
			beego.NSInclude(
				&controllers.ToLowerCaseController{},
			),
		),
		beego.NSNamespace("/docToCode",
			beego.NSInclude(
				&controllers.DocToCodeController{},
			),
		),
		beego.NSNamespace("/tabletostruct",
			beego.NSInclude(
				&controllers.TableToStructController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
