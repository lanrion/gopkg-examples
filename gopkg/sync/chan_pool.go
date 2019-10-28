package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Connection struct {
	ID string
	ConnectAt time.Time
	IsClosed bool
}

type ConPool struct {
	mu  sync.Mutex
	Cap int
	Pool chan *Connection
}

func NewConPool(cap int) (conPool *ConPool) {
	conPool = &ConPool{
		Cap:    cap,
		Pool:   make(chan *Connection, cap),
	}

	for i:=0; i < conPool.Cap; i++ {
		conPool.Pool <- &Connection{ID: getConnectionId(), ConnectAt: time.Now(), IsClosed: false}
	}
	return
}

func (p *ConPool) Get() (con *Connection) {

	for  {
		select {
		case con = <- p.Pool:
			return
		}
	}

}

func (p *ConPool) Put() {
	p.mu.Lock()
	defer p.mu.Unlock()

	select {
	case p.Pool <- &Connection{ID: getConnectionId(), ConnectAt: time.Now(), IsClosed: false}:
		fmt.Println("添加一个连接")
	default:
		fmt.Println("连接池已满")
	}
}

func getConnectionId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
