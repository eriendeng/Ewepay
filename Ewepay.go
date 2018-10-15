package Ewepay

import (
	"github.com/erienniu/Ewepay/conf"
	"github.com/erienniu/Ewepay/order"
)

func Config(path, mode string)  {
	conf.LoadConfig(path).Use(mode)
	return
}

func GetConfig()  {
	conf.GetConfig()
	return
}

func PrepareSubscriptionXML(name, openid string, fee int) (str string, req order.UnifyOrderReq){
	return order.PrepareSubscriptionXML(name, openid, fee)
}

func SubscriptionOrder(str string) {
	order.SubscriptionOrder(str)
}