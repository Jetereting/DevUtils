package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"github.com/astaxie/beego/orm"
	"fmt"
)

// 导出到excel
type MysqlToExcelController struct {
	beego.Controller
}
// @Summary mysql导出到excel
// @Param body body string true "sql内容"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /toExcel [put]
func (c *MysqlToExcelController) ToExcel() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)
	var datas []map[string]interface{}
	o := orm.NewOrm()
	num, err := o.Raw(body).QueryRows(&datas)
	if err == nil {
		fmt.Println("nums: ", num)
		fmt.Println("datas:",datas)
	}else {
		beego.Error(err,body)
	}
}


