package order

import (
	"sync"
	"time"
	"strconv"
	"math/rand"
)

const OrderPostUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"

var mutex sync.Mutex
var channel  = make(chan string)

var OrderAmount = 1

//POST xml结构
type UnifyOrderReq struct {
	Appid            string `xml:"appid"`
	Body             string `xml:"body"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Notify_url       string `xml:"notify_url"`
	Trade_type       string `xml:"trade_type"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee"`
	Out_trade_no     string `xml:"out_trade_no"`
	Sign             string `xml:"sign"`
}

//response xml结构
type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}

//生成唯一订单号
func OrderNumber() string {
	go func() {
		mutex.Lock()
		order := time.Now().Format("20060102150405")+strconv.Itoa(OrderAmount)
		OrderAmount += 1
		channel<-order
		mutex.Unlock()
	}()
	return <-channel
}

func RandomString() string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 30; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}