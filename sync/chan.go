package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 实现简易的队列

type Message struct {
	MsgId int64
	Data  int
}

var queue = make(chan Message, 10)

func addMsg(data int, done chan struct{}) {
	msg := Message{Data: data, MsgId: rand.Int63()}
	queue <- msg

	time.Sleep(time.Second * 2)
	close(done)

}

func readMsg(n int) {
	for {
		select {
		case msg := <-queue:
			fmt.Println("ReadMsg:", n, msg.MsgId, msg.Data)
		}
	}
}

func StartQueue() {
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go readMsg(i)
	}

	addMsg(rand.Int(), done)

	<-done

	// t := time.NewTicker(time.Millisecond * 1)
	// defer t.Stop()
	//
	// for {
	// 	select {
	// 	case <-t.C:
	// 		addMsg(rand.Int())
	// 	}
	// }

}

func Use() {

}
