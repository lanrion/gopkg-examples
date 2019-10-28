package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"reflect"
)

type FakeData struct {
	Category []string
}

func main() {
	name, _ := faker.GetPerson().FirstName(reflect.Value{})
	fmt.Println(name)


}



