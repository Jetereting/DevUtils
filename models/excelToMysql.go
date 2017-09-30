package models

import "github.com/astaxie/beego/orm"

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(ExcelTable))
}

type ExcelTable struct {
	Id      int `orm:"column(id);auto" json:"id"`
	Mark    string
	Uuid    string
	Sheet   string
	Field   string `orm:"column(Field)"`
	Field1  string
	Field2  string
	Field3  string
	Field4  string
	Field5  string
	Field6  string
	Field7  string
	Field8  string
	Field9  string
	Field10 string
	Field11 string
	Field12 string
	Field13 string
	Field14 string
	Field15 string
	Field16 string
	Field17 string
	Field18 string
	Field19 string
	Field20 string
	Field21 string
	Field22 string
	Field23 string
	Field24 string
	Field25 string
	Field26 string
	Field27 string
	Field28 string
	Field29 string
	Field30 string
	Field31 string
	Field32 string
	Field33 string
	Field34 string
	Field35 string
	Field36 string
	Field37 string
	Field38 string
	Field39 string
	Field40 string
	Field41 string
	Field42 string
	Field43 string
	Field44 string
	Field45 string
	Field46 string
	Field47 string
	Field48 string
	Field49 string
	Field50 string
	Field51 string
	Field52 string
	Field53 string
	Field54 string
	Field55 string
	Field56 string
	Field57 string
	Field58 string
	Field59 string
	Field60 string
	Field61 string
}
