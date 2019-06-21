package sync

import (
	"fmt"
	"sync"
)

// Mutex 是一个互斥锁
// 是一种用于多线程编程中，防止两条线程同时对同一公共资源（比如全局变量）进行读写的机制。该目的通过将代码切片成一个一个的临界区域（critical section）达成。临界区域指的是一块对公共资源进行访问的代码，并非一种机制或是算法。一个程序、进程、线程可以拥有多个临界区域，但是并不一定会应用互斥锁。
// 在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
// 使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
// 在 Lock() 之前使用 Unlock() 会导致 panic 异常
// 已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
// 在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
// 适用于读写不确定，并且只有一个读或者写的场景

func TestMutex() {
	var wg sync.WaitGroup

	var mutex sync.Mutex

	fmt.Println("Locked")
	mutex.Lock()
	count := 3

	// 如果不添加lock的话, 打印出来的i值是无序的
	for index := 0; index < count; index++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("not locked", i)
			mutex.Lock()
			fmt.Println("locked", i)
			mutex.Unlock()
			fmt.Println("Unlock", i)
			defer wg.Done()
		}(index)
	}

	mutex.Unlock()
	fmt.Println("Unlock")
	// 必须要在unlock执行后
	wg.Wait()

}
