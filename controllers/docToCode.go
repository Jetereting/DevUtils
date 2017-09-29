package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

// 文档生成代码
type DocToCodeController struct {
	beego.Controller
}

// @Summary 生成apiphone的controller层
// @Param body body string true "页面名称|名称英文+回车+接口名称|名称英文|参数+参数|i（是否带用户码）"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /apiPhoneController [put]
func (c *DocToCodeController) ApiPhoneController() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)

	pageName :=""

	result:=""
	for k, v := range strings.Split(body, "\n") {
		functionName:=""
		args:=""
		if k==0{
			for kk,vv:=range strings.Split(v,"|"){
				if kk==0 {
					result += "//" + vv + "\n"
				}
				if kk==1{
					pageName =vv
					result+="type "+vv+"Controller struct {\n"
					result+="\t"+"controllers.BaseController\n"
				}
			}
		}else {
			for kk,vv:=range strings.Split(v,"|"){
				if kk==0 {
					result += "//" + vv + "\n"
				}else if kk==1{
					functionName=vv
					result+="func (c *"+ pageName +"Controller) "+vv+"() {\n"
				}else if kk==2{
					if len(vv)!=0{
						result+="\tbody, err := c.APIRequestBody()\n"
						result+="\tif err != nil {\n"
						result+="\t\tc.ErrorText(1, \"数据格式错误\")\n"
						result+="\t\treturn\n"
						result+="\t}\n"
						for _,vvv:=range strings.Split(vv,"+"){
							result+="\t"+vvv+", ok := body[\""+vvv+"\"].(string)\n"
							result+="\tif !ok {\n"
							result+="\t\tc.ErrorText(1, \""+vvv+" 错误\")\n"
							result+="\t\treturn\n"
							result+="\t}\n"
							args+=vvv+","
						}
						args=args[:len(args)-1]
					}
				}else if kk==3{
					if len(vv)!=0{
						if len(args)==0{
							args+="intelUserCode"
						}else {
							args += ",intelUserCode"
						}
						result+="\tintelUserCode := c.GetIntelUserCode()\n"
						result+="\tif intelUserCode == \"\" {\n"
						result+="\t\tc.ErrorText(1, \"用户身份错误，请联系管理员\")\n"
						result+="\t\treturn\n"
						result+="\t}\n"
					}
					result+="\th := new(apiphone."+ pageName +"Handle)\n"
					result+="\tc.FeedBack(h."+ functionName +"("+args+"))\n"
				}

			}
		}
		result+="}\n\n"
	}

	c.Ctx.WriteString(result)
}


// @Summary 生成apiphone的handle层
// @Param body body string true "页面名称|名称英文+回车+接口名称|名称英文|参数+参数|i（是否带用户码）"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /apiPhoneHandle [put]
func (c *DocToCodeController) ApiPhoneHandles() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)

	pageName :=""

	result:=""
	for k, v := range strings.Split(body, "\n") {
		args:=""
		if k==0{
			for kk,vv:=range strings.Split(v,"|"){
				if kk==0 {
					result += "//" + vv + "\n"
				}
				if kk==1{
					pageName =vv
					result+="type "+vv+"Handle struct {\n"
					result+="\t"+"handles.BaseHandle\n"
				}
			}
		}else {
			for kk,vv:=range strings.Split(v,"|"){
				if kk==0 {
					result += "//" + vv + "\n"
				}else if kk==1{
					result+="func (h *"+ pageName +"Handle) "+vv+"("
				}else if kk==2{
					if len(vv)!=0{
						for _,vvv:=range strings.Split(vv,"+"){
							args+=vvv+","
						}
						args=args[:len(args)-1]
					}
				}else if kk==3{
					if len(vv)!=0{
						if len(args)==0{
							args+="intelUserCode"
						}else {
							args += ",intelUserCode"
						}
					}

					if len(args)==0{
						result+=args+") models.FeedbackMessage {\n"

						result+="\tsql :=`;`\n"
						result+="\tdata, err := mysqlUtility.DContext.QueryData(sql"+args+")\n"
					}else {
						result+=args+" string) models.FeedbackMessage {\n"

						result+="\tsql :=`;`\n"
						result+="\tdata, err := mysqlUtility.DContext.QueryData(sql, "+args+")\n"
					}
					result+="\tif err != nil {\n"
					result+="\t\treturn h.Feedback(1, \"获取错误\" + err.Error(), nil)\n"
					result+="\t}\n"
					result+="\tif len(data) == 0 {\n"
					result+="\t\treturn h.Feedback(1, \"没有数据\", nil)\n"
					result+="\t}\n"
					result+="\treturn h.Feedback(0, \"成功\", data)\n"
				}
			}
		}
		result+="}\n\n"
	}

	c.Ctx.WriteString(result)
}
