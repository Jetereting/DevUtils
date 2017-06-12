package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

// model 和 markdown 的相互转化
type ModelTranslateMarkDownController struct {
	beego.Controller
}

// @Summary model to markdown
// @Param body body string true "结构体里面的内容,此接口只支持bson,按照： |属性|类型|数据库字段|备注|"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /ModelToMarkDown [put]
func (c *ModelTranslateMarkDownController) ModelToMarkDown() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)
	body = strings.Replace(body, "\"", "", -1)
	body = strings.Replace(body, "`", "", -1)
	body = strings.Replace(body, ":", "", -1)
	body = strings.Replace(body, "//", " ", -1)
	body = strings.Replace(body, "bson", "", -1)
	body = strings.Replace(body, "description", "", -1)

	result :=
		`|属性|类型|数据库字段|备注|
|---|---|----|----|` + "\n"

	for _, v := range strings.Split(body, "\n") {
		for _, vv := range strings.Split(v, " ") {
			if len(vv) == 0 {
				continue
			}
			result += `|` + vv
		}
		result += `|` + "\n"
	}
	result = strings.Replace(result, "	", "", -1)
	c.Ctx.WriteString(result)
}

// @Summary markdown to model
// @Param body body string true "markdown里面的内容,此接口只支持bson,按照： |属性|类型|数据库字段|备注|"
// @Param description query bool false "备注是否放在描述里面"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /MarkDownToModel [put]
func (c *ModelTranslateMarkDownController) MarkDownToModel() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)
	description, _ := c.GetBool("description", false)
	result := ""
	for k, v := range strings.Split(body, "\n") {
		if k == 0 || k == 1 {
			continue
		}
		for kk, vv := range strings.Split(v, "|") {
			if description {
				if kk == 4 {
					result += vv + "\"`    "
				} else if kk == 3 {
					result += "`bson:\"" + vv + "\" description:\""
				} else {
					result += vv + "    "
				}
			} else {
				if kk == 4 {
					result += "//" + vv
				} else if kk == 3 {
					result += "`bson:\"" + vv + "\"`    "
				} else {
					result += vv + "    "
				}
			}
		}
		result += "\n"
	}
	c.Ctx.WriteString(result)
}
