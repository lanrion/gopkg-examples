package main

import (
	"github.com/go-redis/redis"
	"log"
)

func Sub1() {
	log.Println("Sub1...")
	sub := GetCache().Subscribe("testUser")
	defer sub.Close()
	for {
		iface, _ := sub.Receive()
		switch m := iface.(type) {
		case *redis.Subscription:
			log.Printf("Subscription1 Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
		case *redis.Message:
			log.Printf("Message1: %v", m.String())

			// received first message
		case *redis.Pong:
			// pong received
		default:
			// handle error
		}
	}
}

func Sub2() {
	log.Println("sub2s...")
	sub := GetCache().Subscribe("testUser")
	defer sub.Close()
	ch := sub.Channel()

	for {
		msg, ok := <-ch
		if !ok {
			continue
		}
		log.Printf("Sub3: %v", msg.String())
	}
}

func main() {
	Sub1()
}
