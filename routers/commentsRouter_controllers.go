package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getTempTableSql`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"],
		beego.ControllerComments{
			Method: "ParseFile",
			Router: `/parseFile/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
