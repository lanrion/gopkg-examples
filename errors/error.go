package errors

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

// recover 只允许在defer函数中使用
// panic相当于直接抛出异常 raise

func LogErr()(err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("err: %+v", r)
			log.Println("work failed:", err)
		}

		if err != nil {
			log.Println("work failed11:", err)
		}
	}()

	err = errors.New("1234321")
	return
}


func TestError() {

	var wg sync.WaitGroup

	wg.Add(1)
	go testArrayIndex(&wg)
	wg.Wait()

	// panic相当于直接抛出异常 raise
	// panic("no value for $USER")
}

func testArrayIndex(wg *sync.WaitGroup) {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	defer wg.Done()
	defer logErr()

	for index := 0; index < 8; index++ {
		val := arr[index]
		fmt.Println(val)
	}
}

func logErr() {
	if err := recover(); err != nil {
		log.Println("work failed:", err)
	}
}
