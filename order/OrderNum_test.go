package order_test

import (
	"testing"
	"fmt"
	"../order"
)

func TestOrderNumber(t *testing.T) {
	var i int
	for i=0; i<10000; i++ {
		fmt.Println(order.OrderNumber())
	}
}
