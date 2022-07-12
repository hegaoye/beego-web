package job

import (
	"com.demo/webdemo/rpcclient"
	"com.demo/webdemo/service"
	"context"
	"fmt"
	"github.com/beego/beego/v2/task"
)

var userService *service.UserTaskService

func init() {
	userService = new(service.UserTaskService)
}

func Job() {
	rpcTestTask := task.NewTask("rpcTestTask", "* * * * * *", rpcTestTask)
	updateTask := task.NewTask("updateTask", "* * * * * *", updateTask)
	uploadTask := task.NewTask("uploadTask", "0 0 0 * * *", uploadTask)
	heartbeatTask := task.NewTask("heartbeatTask", "0 30 * * * *", heartbeatTask)

	task.AddTask("rpcTestTask", rpcTestTask)
	task.AddTask("heartbeatTask", heartbeatTask)
	task.AddTask("updateTask", updateTask)
	task.AddTask("uploadTask", uploadTask)
	task.StartTask()

	//defer task.StopTask()

}

func rpcTestTask(ctx context.Context) error {
	fmt.Println("rpc test......")
	rpcclient.SendToServerMessageResult("this text from task job")
	return nil
}

// 心跳
func heartbeatTask(ctx context.Context) error {
	//userService.Heartbeat()
	return nil
}

// 上傳任務
func uploadTask(ctx context.Context) error {
	//userService.Upload()
	return nil
}

// 更新任務
func updateTask(ctx context.Context) error {
	//userService.UpdateDemo()
	return nil
}
