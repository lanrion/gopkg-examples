package data

import (
	"fmt"
	"testing"
)

func TestArray1(t *testing.T) {
	Array1()
}

func TestGroupOrder(t *testing.T) {
	GroupOrder(100 * 1000)
}

func TestGenerateOrders(t *testing.T) {
	arr := GenerateOrders(100)
	fmt.Println("Len: ", len(arr))
}

func BenchmarkGenerateOrders(b *testing.B) {
	arr := GenerateOrders(100)
	fmt.Println("Len: ", len(arr))
}

func TestEverything(t *testing.T) {
	fs := map[string]func(*testing.T){"testFoo": TestArray1, "testBar": TestGroupOrder, "testBaz": TestGenerateOrders}
	for name, f := range fs {
		t.Run(name, f)
	}
}
