package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Sex  int
}

func (m *Student) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Student) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type Student struct {
	Age int
	User
}

func testRedis() {

	arrays := []string{"Name", "Age", "Sex"}
	data, _ := json.Marshal(arrays)

	_, err := GetCache().Set("arrays", data, 0).Result()

	fmt.Printf(" %v, %v", err, data)

	str, err := GetCache().Get("arrays").Bytes()

	var a1 []string
	err = json.Unmarshal(str, &a1)

	// fmt.Printf("string, %v, %v, %v", err, str, a1)

	fmt.Println(strconv.Itoa(12))
}

func Client() {

}
