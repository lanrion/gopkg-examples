package container

import (
	"container/list"
	"fmt"
)

// 双向链表
// https://golang.google.cn/pkg/container/list

func TestList() {
	l := list.New()
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack("123")

	// 获取前面的第一个, 然后循环操作
	fmt.Println("从前打印:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("从后打印:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// 重置list
	l.Init()
	fmt.Println(l.Len())
}
