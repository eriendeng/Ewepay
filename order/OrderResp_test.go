package order_test

import "testing"
import (
	"../order"
	"../conf"
	"fmt"
)

func TestSubscriptionOrder(t *testing.T) {
	conf.LoadConfig("/Users/erien/go/myprojects/Ewepay/.properties").Use("DEVELOPMENT")
	str, _ := order.PrepareSubscriptionXML("test",1)
	fmt.Println(str)
	order.SubscriptionOrder(str)
}