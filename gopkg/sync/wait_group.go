package main

import (
	"fmt"
	"runtime/pprof"
	"sync"
)

// WaitGroup主要是等待多个goroutine执行结束
// 通过计算器的形式来判断是否全部执行完
// 在main的routine中Add设置goroutine的总数, 同时Wait
// 在所有routing中进行Done
// Add:  如果是0, 阻塞会全部释放, 如果是负数, 则会抛出异常
// Done就是, 计数器-1的操作
// Wait: 启动一个for循环等待, 计数器为0时, 则表示可以退出main函数
// 计算器是通过state来维护
// https://golang.org/src/sync/waitgroup.go

func TestWaitGroup() {
	var wg sync.WaitGroup
	// 如果将3改为<3的值, 那打印出来的只有<3的结果
	testChan := make(chan int)
	wg.Add(3)
	count := 3
	for index := 0; index < count; index++ {
		go func(i int) {
			defer wg.Done()
			testChan <- i
			// 这里打印出来的是没有顺序的
			fmt.Println(i)
		}(index)
	}

	go func() {
		wg.Wait()
		close(testChan)
	}()

	for v := range testChan {
		fmt.Println("V: ", v)
	}

	// wg.Wait()
	// close(testChan)

}

func GetRoutineCount() {
	//profile := pprof.Lookup("goroutine")
	wg := sync.WaitGroup{}
	//for i := 0; i < d.cfg.Concurrency; i++ {
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			fmt.Println("parse err: ")
		}()
	}

	profile := pprof.Lookup("goroutine")

	fmt.Println(profile.Count())
}
