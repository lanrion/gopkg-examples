package main

import (
	"fmt"
	"sync"
	"time"
)

// 用于控制多个goroutine之间的执行
// 当某个条件变量发生变化时, 然后通知Wait方法继续执行

// 条件变量的作用并不是保证在同一时刻仅有一个线程访问某一个共享数据，
// 而是在对应的共享数据的状态发生变化时，通知其他因此而被阻塞的线程。条件变量总是与互斥量组合使用。
// 互斥量为共享数据的访问提供互斥支持，而条件变量可以就共享数据的状态的变化向相关线程发出通知。

// TestCond, 只有两个房间可供使用, 需要保证每个有人退房时, 马上被入住
func TestCond() {
	cond := sync.NewCond(new(sync.Mutex))
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue")
		cond.L.Unlock()
		cond.Signal()
	}

	// 处理10个客人的入住需求
	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(queue) == 2 {
			// 当房间满了, 要等待退房的removeFromQueue Signal/Broadcast通知
			cond.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		cond.L.Unlock()
	}

	fmt.Printf("queue len %v \n", len(queue))

}
