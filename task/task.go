package task

import (
	"github.com/astaxie/beego/toolbox"
	"fmt"
)

func StartTask()  {
	tk1 := toolbox.NewTask("Statistics", "0/3 * * * * *", func() error{
		fmt.Println("lalallaaaaaa")
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
}
func StopTask() {
	toolbox.StopTask()
}