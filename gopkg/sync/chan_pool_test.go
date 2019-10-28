package main

import (
	"fmt"
	"testing"
)

func TestNewConPool(t *testing.T) {
	ncp := NewConPool(2)
	ncp.Get()
	ncp.Put()
	fmt.Println(len(ncp.Pool))


}
