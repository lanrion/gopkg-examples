package data

import (
	"fmt"
)

// https://gobyexample.com/arrays
// 基本的操作
func TestArray() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set: ", a)

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dcl: ", b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
