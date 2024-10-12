package robfigcron01

import (
	"fmt"

	"github.com/robfig/cron"
)

func RobfigCron01() {
	// 创建一个新的 cron 调度器
	c := cron.New()

	// 添加一个定时任务，每分钟执行一次
	err := c.AddFunc("* * * * *", func() {
		// 定时任务 业务逻辑...

	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	// 启动 cron 调度器
	c.Start()

	// 阻塞主线程，让定时任务一直执行
	select {}
}
