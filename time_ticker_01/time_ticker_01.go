package timeticker01

import (
	"fmt"
	"time"
)

// 创建一个定时器，每隔1秒触发一次, 5秒后停止定时器
func TimeTicker01() {
	ticker := time.NewTicker(1 * time.Second)

	// 创建一个退出信号通道
	quit := make(chan struct{})

	// 启动一个goroutine来处理定时器事件
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("定时器触发")
				// ... 处理定时器事件的逻辑 ...

			case <-quit:
				// 收到退出信号，停止定时器
				ticker.Stop()
				return
			}
		}
	}()

	// 模拟运行5秒后停止定时器
	time.Sleep(5 * time.Second)
	quit <- struct{}{}
	fmt.Println("定时器已停止")
}
