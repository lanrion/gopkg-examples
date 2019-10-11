package main

import (
	"fmt"
	"os/exec"
)

type Animal struct {
	Name string
}

func main() {

	str, err :=exec.LookPath("mogri111fy")
	fmt.Println(err, str)

}

func test1(dog interface{}) {
	mydog := dog.(Animal)
	fmt.Println(mydog.Name)
}


