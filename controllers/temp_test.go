package controllers

import (
	"testing"
	"fmt"
	"strings"
)

func TestTemp(t *testing.T) {
	result:=""
	a:="迎新报到|YxCheckIn"
	b:=strings.Split(a,"|")
	for k,v:=range b{
		if k==0 {
			result += "//" + v + "\n"
		}
		if k==1{
			result+=`type `+v+`Controller struct {
	controllers.BaseController
}`
		}
	}
	fmt.Println(result)
}
