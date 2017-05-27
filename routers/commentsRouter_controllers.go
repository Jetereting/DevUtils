package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:MainController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/a/getTempTableSql`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:MainController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:MainController"],
		beego.ControllerComments{
			Method: "UploadExcel",
			Router: `/b/file/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beegoAutoDoc/controllers:MainController"] = append(beego.GlobalControllerRouter["beegoAutoDoc/controllers:MainController"],
		beego.ControllerComments{
			Method: "ParseExcel",
			Router: `/c/file/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
