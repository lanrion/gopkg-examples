package data

import (
	"fmt"
	"strconv"
)

func CovertToInt(string2 string) {
	int1, _ := strconv.Atoi(string2)
	fmt.Println("CovertToInt: ", int1)
}