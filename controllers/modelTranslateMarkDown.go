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
// @Param is3 query bool false "|字段|类型|说明|"
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
// @Param body body string true "markdown里面的内容,此接口只支持bson,按照： |属性|类型|字段|备注|  或 |字段|类型|说明| "
// @Param description query bool false "备注是否放在描述里面"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /MarkDownToModel [put]
func (c *ModelTranslateMarkDownController) MarkDownToModel() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)

	body = strings.Replace(body, "字符串数组", "[]string", -1)
	body = strings.Replace(body, "数值数组", "[]int", -1)
	body = strings.Replace(body, "字符串", "string", -1)
	body = strings.Replace(body, "数值", "int", -1)
	body = strings.Replace(body, "整数", "int", -1)
	body = strings.Replace(body, " ", "", -1)

	description, _ := c.GetBool("description", false)
	result := ""
	for k, v := range strings.Split(body, "\n") {
		if k == 0 || k == 1 {
			continue
		}
		num := len(strings.Split(v, "|"))
		vvs := strings.Split(v, "|")
		switch num {
		case 5:
			if description {
				result += vvs[1] + "    "
				result += vvs[2] + "    "
				result += "`bson:\"" + vvs[1] + "\""
				if len(vvs[3])==0{
					result += "`    "
				}else {
					result += " description:\"" + vvs[3] + "\"`    "
				}
			}else {
				result += vvs[1] + "    "
				result += vvs[2] + "    "
				result += "`bson:\"" + vvs[1] + "\"`    "
				if len(vvs[3])!=0{
					result += "//" + vvs[3]
				}
			}
			break
		case 6:
			if description {
				result += vvs[2] + "    "
				result += vvs[3] + "    "
				result += "`bson:\"" + vvs[2] + "\""
				if len(vvs[4])==0{
					result += "`    "
				}else {
					result += " description:\"" + vvs[4] + "\"`    "
				}
			}else {
				result += vvs[2] + "    "
				result += vvs[3] + "    "
				result += "`bson:\"" + vvs[2] + "\"`    "
				if len(vvs[4])!=0{
					result += "//" + vvs[4]
				}
			}
			break
		}
		result += "\n"
	}
	c.Ctx.WriteString(result)
}
