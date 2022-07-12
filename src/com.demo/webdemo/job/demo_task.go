package job

import (
	"com.demo/webdemo/service"
	"context"
	"github.com/beego/beego/v2/task"
)

var userService *service.UserTaskService

func init() {
	userService = new(service.UserTaskService)
}

func Job() {
	updateTask := task.NewTask("updateTask", "* * * * * *", updateTask)
	uploadTask := task.NewTask("uploadTask", "0 0 0 * * *", uploadTask)
	heartbeatTask := task.NewTask("heartbeatTask", "0 30 * * * *", heartbeatTask)

	task.AddTask("heartbeatTask", heartbeatTask)
	task.AddTask("updateTask", updateTask)
	task.AddTask("uploadTask", uploadTask)
	task.StartTask()

	//defer task.StopTask()

}

// 心跳
func heartbeatTask(ctx context.Context) error {
	userService.Heartbeat()
	return nil
}

// 上傳任務
func uploadTask(ctx context.Context) error {
	userService.Upload()
	return nil
}

// 更新任務
func updateTask(ctx context.Context) error {
	//userService.UpdateDemo()
	return nil
}
