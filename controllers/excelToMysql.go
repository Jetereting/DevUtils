package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"github.com/tealeg/xlsx"
	"github.com/astaxie/beego/orm"
	"beegoAutoDoc/models"
	"os"
	"strings"
)

// excle导入到临时表
type ExcelToMysqlController struct {
	beego.Controller
}

// @Summary 获取创建temp_table语句
// @Title 获取创建temp_table语句
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /getTempTableSql [get]
func (c *ExcelToMysqlController) Get() {
	result := make(map[string]interface{})
	result["text"] =
		`
		CREATE TABLE qdxg.temp_table
		(
		  Id INT(11) PRIMARY KEY NOT NULL COMMENT 'id' AUTO_INCREMENT,
		  Mark VARCHAR(50),
		  Field   VARCHAR(50),
		  Field1  VARCHAR(50),
		  Field2  VARCHAR(50),
		  Field3  VARCHAR(50),
		  Field4  VARCHAR(50),
		  Field5  VARCHAR(50),
		  Field6  VARCHAR(50),
		  Field7  VARCHAR(50),
		  Field8  VARCHAR(50),
		  Field9  VARCHAR(50),
		  Field10 VARCHAR(50),
		  Field11 VARCHAR(50),
		  Field12 VARCHAR(50),
		  Field13 VARCHAR(50),
		  Field14 VARCHAR(50),
		  Field15 VARCHAR(50),
		  Field16 VARCHAR(50),
		  Field17 VARCHAR(50),
		  Field18 VARCHAR(50),
		  Field19 VARCHAR(50),
		  Field20 VARCHAR(50),
		  Field21 VARCHAR(50),
		  Field22 VARCHAR(50),
		  Field23 VARCHAR(50),
		  Field24 VARCHAR(50),
		  Field25 VARCHAR(50),
		  Field26 VARCHAR(50),
		  Field27 VARCHAR(50),
		  Field28 VARCHAR(50),
		  Field29 VARCHAR(50),
		  Field30 VARCHAR(50),
		  Field31 VARCHAR(50),
		  Field32 VARCHAR(50),
		  Field33 VARCHAR(50),
		  Field34 VARCHAR(50),
		  Field35 VARCHAR(50),
		  Field36 VARCHAR(50),
		  Field37 VARCHAR(50),
		  Field38 VARCHAR(50),
		  Field39 VARCHAR(50),
		  Field40 VARCHAR(50),
		  Field41 VARCHAR(50),
		  Field42 VARCHAR(50),
		  Field43 VARCHAR(50),
		  Field44 VARCHAR(50),
		  Field45 VARCHAR(50),
		  Field46 VARCHAR(50),
		  Field47 VARCHAR(50),
		  Field48 VARCHAR(50),
		  Field49 VARCHAR(50),
		  Field50 VARCHAR(50),
		  Field51 VARCHAR(50),
		  Field52 VARCHAR(50),
		  Field53 VARCHAR(50),
		  Field54 VARCHAR(50),
		  Field55 VARCHAR(50),
		  Field56 VARCHAR(50),
		  Field57 VARCHAR(50),
		  Field58 VARCHAR(50),
		  Field59 VARCHAR(50),
		  Field60 VARCHAR(50),
		  Field61 VARCHAR(50)
		);


		`
	c.Ctx.WriteString(result["text"].(string))
}


// @Summary 上传xlsx文件并解析到临时表
// @Title 上传xlsx文件并解析到临时表
// @Param fileData formData file true "文件数据(暂时只支持xlsx格式的 且只有一个sheet,小于62列) 导入完成后该文件会被自动删除"
// @Param mark query string true "用于区分导入的文件，对应temp_table表的 Mark 字段"
// @Success 200 成功
// @Failure 400 请求发生错误
// @Failure 500 服务器错误
// @router /parseFile/ [post]
func (c *ExcelToMysqlController) ParseFile() {
	filePath:=""
	f, h, err := c.GetFile("fileData")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
		c.Ctx.WriteString("错误："+err.Error())
		return
	} else {
		filePath = "static/upload/" + time.Now().Format("20060102150405") + h.Filename
		err = c.SaveToFile("fileData", filePath)
		if err!=nil{
			c.Ctx.WriteString("上传保存失败："+err.Error())
			return
		}
	}

	if strings.Split(filePath,".")[1]!="xlsx"{
		c.Ctx.WriteString("暂时只支持xlsx格式的 且只有一个sheet："+err.Error())
		return
	}
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		c.Ctx.WriteString("打开xlsx文件错误，你是不是用其他软件处理了？可以再处理回来的："+err.Error())
		return
	}
	if len(file.Sheets) == 0 {
		c.Ctx.WriteString("excel中没有数据，请核实："+err.Error())
		return
	}

	sheet := file.Sheets[0]
	var tempTables []models.TempTable
	for i, row := range sheet.Rows {
		if i < 1 {
			continue
		}
		tempTable := models.TempTable{}
		for k, v := range row.Cells {
			tempTable.Mark=c.GetString("mark")
			switch k {
			case 0:tempTable.Field = v.Value
			case 1:tempTable.Field1 = v.Value
			case 2:tempTable.Field2 = v.Value
			case 3:tempTable.Field3 = v.Value
			case 4:tempTable.Field4 = v.Value
			case 5:tempTable.Field5 = v.Value
			case 6:tempTable.Field6 = v.Value
			case 7:tempTable.Field7 = v.Value
			case 8:tempTable.Field8 = v.Value
			case 9:tempTable.Field9 = v.Value
			case 10:tempTable.Field10 = v.Value
			case 11:tempTable.Field11 = v.Value
			case 12:tempTable.Field12 = v.Value
			case 13:tempTable.Field13 = v.Value
			case 14:tempTable.Field14 = v.Value
			case 15:tempTable.Field15 = v.Value
			case 16:tempTable.Field16 = v.Value
			case 17:tempTable.Field17 = v.Value
			case 18:tempTable.Field18 = v.Value
			case 19:tempTable.Field19 = v.Value
			case 20:tempTable.Field20 = v.Value
			case 21:tempTable.Field21 = v.Value
			case 22:tempTable.Field22 = v.Value
			case 23:tempTable.Field23 = v.Value
			case 24:tempTable.Field24 = v.Value
			case 25:tempTable.Field25 = v.Value
			case 26:tempTable.Field26 = v.Value
			case 27:tempTable.Field27 = v.Value
			case 28:tempTable.Field28 = v.Value
			case 29:tempTable.Field29 = v.Value
			case 30:tempTable.Field30 = v.Value
			case 31:tempTable.Field31 = v.Value
			case 32:tempTable.Field32 = v.Value
			case 33:tempTable.Field33 = v.Value
			case 34:tempTable.Field34 = v.Value
			case 35:tempTable.Field35 = v.Value
			case 36:tempTable.Field36 = v.Value
			case 37:tempTable.Field37 = v.Value
			case 38:tempTable.Field38 = v.Value
			case 39:tempTable.Field39 = v.Value
			case 40:tempTable.Field40 = v.Value
			case 41:tempTable.Field41 = v.Value
			case 42:tempTable.Field42 = v.Value
			case 43:tempTable.Field43 = v.Value
			case 44:tempTable.Field44 = v.Value
			case 45:tempTable.Field45 = v.Value
			case 46:tempTable.Field46 = v.Value
			case 47:tempTable.Field47 = v.Value
			case 48:tempTable.Field48 = v.Value
			case 49:tempTable.Field49 = v.Value
			case 50:tempTable.Field50 = v.Value
			case 51:tempTable.Field51 = v.Value
			case 52:tempTable.Field52 = v.Value
			case 53:tempTable.Field53 = v.Value
			case 54:tempTable.Field54 = v.Value
			case 55:tempTable.Field55 = v.Value
			case 56:tempTable.Field56 = v.Value
			case 57:tempTable.Field57 = v.Value
			case 58:tempTable.Field58 = v.Value
			case 59:tempTable.Field59 = v.Value
			case 60:tempTable.Field60 = v.Value
			case 61:tempTable.Field61 = v.Value
			}
		}
		tempTables=append(tempTables,tempTable)
	}
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, tempTables)
	if err!=nil{
		c.Ctx.WriteString("插入失败："+err.Error())
	}
	c.Ctx.WriteString(fmt.Sprint("成功！插入行数：",successNums))
	defer os.Remove(filePath)
}
