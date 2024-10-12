package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func CpuWork() {

	runtime.GOMAXPROCS(1)

	go StartPprof(":6060")

	CpuWorkload()

	select {}
}

func StartPprof(port string) {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		println("启动 pprof 错误", err)
	}
}

// 模拟一些 CPU 工作负载
func CpuWorkload() {
	var mut sync.Mutex
	for i := 0; i < 10; i++ {
		go CpuTaskGroup(mut)
	}
}

func CpuTaskGroup(mut sync.Mutex) {

	for {
		mut.Lock()
		CpuIntensiveTask1()
		mut.Unlock()
		CpuIntensiveTask2()
		time.Sleep(1 * time.Second)
	}

}

// CPU 密集型任务
func CpuIntensiveTask1() {
	i := 0
	for {
		i = i + 1
		if i > 1e9 {
			break
		}
		_ = i
	}
}
func CpuIntensiveTask2() {
	i := 0
	for {
		i = i + 1
		if i > 1e9 {
			break
		}
		_ = i
	}
}
