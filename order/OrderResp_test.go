package order_test

import "testing"
import (
	"github.com/erienniu/Ewepay/order"
	"github.com/erienniu/Ewepay/conf"
	"fmt"
)

func TestSubscriptionOrder(t *testing.T) {
	conf.LoadConfig("/Users/erien/go/myprojects/Ewepay/.properties").Use("DEVELOPMENT")
	str, _ := order.PrepareSubscriptionXML("test","111", 1)
	fmt.Println(str)
	order.SubscriptionOrder(str)
}