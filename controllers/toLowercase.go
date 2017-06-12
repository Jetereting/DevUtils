package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)


// 路由等全部转化成小写
type ToLowerCaseController struct {
	beego.Controller
}

// @Summary 转化为小写
// @Param body body string true "要变成小写的内容"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /toLowerCase [put]
func (c *ToLowerCaseController) ToLowerCase(){
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)
	body=strings.ToLower(body)
	c.Ctx.WriteString(body)
}