package main

import (
	"fmt"
	"testing"
)

func TestGetRoutineCount(t *testing.T) {
	GetRoutineCount()
}

func TestTestWaitGroup(t *testing.T) {
	TestWaitGroup()
}

func TestStartQueue(t *testing.T) {
	StartQueue()
}

func TestRwMutex(t *testing.T) {
	RwMutex()
}

func TestChan1(t *testing.T)  {
	chan1 := make(chan interface{})
	fmt.Println("TestChan1")
	go func() {
		chan1 <- "123"
	}()
	<- chan1
}
