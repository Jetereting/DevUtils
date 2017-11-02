package controllers

import (
	"io/ioutil"
	"github.com/astaxie/beego"
	"strings"
	"beegoAutoDoc/utils"
)


// mysql表生成结构体
type TableToStructController struct {
	beego.Controller
}

// @Summary mysql表生成结构体
// @Param body body string true "mysql 表的创建语句"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /tabletostruct [put]
func (c *TableToStructController) C() {
	req := c.Ctx.Request.Body
	bod, _ := ioutil.ReadAll(req)
	body := string(bod)
	result:=""
	for _,v:=range strings.Split(body,"\n"){
		if strings.Contains(v,"CREATE"){
			result+="type "+utils.GetBetweenStr(v,"`","`")+" struct{\n"
		}
		if strings.Contains(v,"int("){
			field:=utils.GetBetweenStr(v,"`","`")
			result+="\t"+field+" int `db:\""+field+"\"`"
			if strings.Contains(v,"COMMENT"){
				result+=" //"+utils.GetBetweenStr(v,"COMMENT '","'")
			}
			result+="\n"
		}
		if strings.Contains(v,"varchar("){
			field:=utils.GetBetweenStr(v,"`","`")
			result+="\t"+field+" string `db:\""+field+"\"`"
			if strings.Contains(v,"COMMENT"){
				result+=" //"+utils.GetBetweenStr(v,"COMMENT '","'")
			}
			result+="\n"
		}
		if strings.Contains(v,"timestamp"){
			field:=utils.GetBetweenStr(v,"`","`")
			result+="\t"+field+" string `db:\""+field+"\"`"
			if strings.Contains(v,"COMMENT"){
				result+=" //"+utils.GetBetweenStr(v,"COMMENT '","'")
			}
			result+="\n"
		}
		if strings.Contains(v," float("){
			field:=utils.GetBetweenStr(v,"`","`")
			result+="\t"+field+" float64 `db:\""+field+"\"`"
			if strings.Contains(v,"COMMENT"){
				result+=" //"+utils.GetBetweenStr(v,"COMMENT '","'")
			}
			result+="\n"
		}
	}
	result+="}"
	c.Ctx.WriteString(result)
}
