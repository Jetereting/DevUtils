package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:DocToCodeController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:DocToCodeController"],
		beego.ControllerComments{
			Method: "ApiPhoneController",
			Router: `/apiPhoneController`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:DocToCodeController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:DocToCodeController"],
		beego.ControllerComments{
			Method: "ApiPhoneHandles",
			Router: `/apiPhoneHandle`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getTempTableSql`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ExcelToMysqlController"],
		beego.ControllerComments{
			Method: "ParseFile",
			Router: `/parseFile`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ModelTranslateMarkDownController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ModelTranslateMarkDownController"],
		beego.ControllerComments{
			Method: "ModelToMarkDown",
			Router: `/ModelToMarkDown`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ModelTranslateMarkDownController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ModelTranslateMarkDownController"],
		beego.ControllerComments{
			Method: "MarkDownToModel",
			Router: `/MarkDownToModel`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:ToLowerCaseController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:ToLowerCaseController"],
		beego.ControllerComments{
			Method: "ToLowerCase",
			Router: `/toLowerCase`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
