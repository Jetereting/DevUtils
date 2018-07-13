package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["DevUtils/controllers:ConfigController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ConfigController"],
		beego.ControllerComments{
			Method: "MysqlDetail",
			Router: `/mysql`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:DocToCodeController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:DocToCodeController"],
		beego.ControllerComments{
			Method: "ApiPhoneController",
			Router: `/apiPhoneController`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:DocToCodeController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:DocToCodeController"],
		beego.ControllerComments{
			Method: "ApiPhoneHandles",
			Router: `/apiPhoneHandle`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:ExcelToMysqlController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ExcelToMysqlController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getTempTableSql`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:ExcelToMysqlController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ExcelToMysqlController"],
		beego.ControllerComments{
			Method: "ParseFile",
			Router: `/parseFile`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:ModelTranslateMarkDownController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ModelTranslateMarkDownController"],
		beego.ControllerComments{
			Method: "ModelToMarkDown",
			Router: `/ModelToMarkDown`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:ModelTranslateMarkDownController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ModelTranslateMarkDownController"],
		beego.ControllerComments{
			Method: "SchemaToMarkDown",
			Router: `/SchemaToMarkDown`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:ModelTranslateMarkDownController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ModelTranslateMarkDownController"],
		beego.ControllerComments{
			Method: "MarkDownToModel",
			Router: `/MarkDownToModel`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:MysqlToExcelController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:MysqlToExcelController"],
		beego.ControllerComments{
			Method: "ToExcel",
			Router: `/toExcel`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:TableToStructController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:TableToStructController"],
		beego.ControllerComments{
			Method: "C",
			Router: `/tabletostruct`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DevUtils/controllers:ToLowerCaseController"] = append(beego.GlobalControllerRouter["DevUtils/controllers:ToLowerCaseController"],
		beego.ControllerComments{
			Method: "ToLowerCase",
			Router: `/toLowerCase`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
