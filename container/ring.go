package container

import (
	"container/ring"
	"fmt"
)

// 环
// 结构有点特殊，环的尾部就是头部，所以每个元素实际上就可以代表自身的这个环。 它不需要像list一样保持list和element两个结构，只需要保持一个结构就行。

func TestRing() {
	r := ring.New(5)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
