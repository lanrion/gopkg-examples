package data

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// https://gobyexample.com/arrays
// 基本的操作
func Array1() {
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

type Order struct {
	Id int
	Name string
	CreatedAt time.Time
}

func filter(arr []Order, f func(Order) bool)[]Order {
	vsf := make([]Order, 0)
	for _, v := range arr {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func GenerateOrders(num int) (orders []Order) {
	var wg sync.WaitGroup
	baseNum := 10
	goroutineNums := num / baseNum
	orderChan := make(chan Order)
	for i:= 0; i < goroutineNums ; i++ {
		wg.Add(1)
		go func(total int) {
			defer wg.Done()
			for i := 0 ; i < total; i++ {
				name := fmt.Sprintf("OrderName-%d", i)
				createdAt := time.Now().Add(1 * time.Minute)
				order := Order{Id: i, Name: name, CreatedAt: createdAt }
				orderChan <- order
			}
		}(baseNum)
	}

	io.Pipe()

	http.NewRequest()

	io.Copy()

	// 不会死锁
	go func() {
		wg.Wait()
		close(orderChan)
	}()

	// 会死锁

	// wg.Wait()
	// close(orderChan)

	for v := range orderChan {
		orders = append(orders, v)
	}

	return orders
}

func GroupOrder(num int) {
	// arr := GenerateOrders(num)
	// evenOrders := filter(arr, func(order Order) bool {
	// 	return order.Id % 2 == 0
	// })

	// for _, v := range evenOrders {
	// 	fmt.Println("Order: ", v.Name)
	// }
}

