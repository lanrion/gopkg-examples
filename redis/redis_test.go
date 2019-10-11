package main

import (
	"fmt"
	"testing"
	"time"
)

func ConcurrentIncry() {
	res := GetCache().Incr("myk").Val()
	if res <= 5 {
		fmt.Println("TestConcurrentIncry", res, time.Now().Round(time.Millisecond))
	} else {
		fmt.Println("TestConcurrentIncry NONE", res, time.Now().Round(time.Millisecond))

	}
}

func TestConcurrentIncry(t *testing.T) {
	for i := 0; i < 50; i++ {
		go ConcurrentIncry()
	}
	time.Sleep(10 * time.Second)
}
