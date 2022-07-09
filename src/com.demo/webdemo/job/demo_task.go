package job

import (
	"context"
	"fmt"
	"github.com/beego/beego/task"
	"time"
)

func Job() {
	tk1 := task.NewTask("tk1", "* * * * * *", generateWarning)
	task.AddTask("tk1", tk1)
	task.StartTask()
	//defer task.StopTask()

}

func generateWarning(context.Context) error {
	now := time.Now()
	fmt.Println(now.Local())
	return nil
}
