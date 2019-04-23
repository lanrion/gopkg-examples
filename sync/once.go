package sync

import (
	"fmt"
	"sync"
)

// Once, 确保只执行一次
// 适合在配置上使用

func TestOnce() {
	count := 5

	bol := make(chan bool)

	var arr = [5]bool{true, false, false, true, false}

	for index := 0; index < count; index++ {
		go func(i int) {
			bol <- arr[i]
		}(index)
	}
	for index := 0; index < count; index++ {
		fmt.Println(<-bol)
	}

	fmt.Println("使用Once后")
	// 使用Once后
	var once sync.Once
	for index := 0; index < count; index++ {
		go func(i int) {
			once.Do(func() {
				bol <- arr[i]
			})
		}(index)
	}

	count = 1
	// 如果将count改为>1的, 会报错, 因为只塞入了一个chan
	for index := 0; index < count; index++ {
		fmt.Println(<-bol)
	}

}
