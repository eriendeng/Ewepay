# Ewepay
##### 基于golang的微信开发平台整合
--

![微信支付流程](https://pay.weixin.qq.com/wiki/doc/api/img/chapter7_4_1.png)
基于微信支付流程，整合大量琐碎结构和接口。

* 支持configuration中写入所需信息
* 写入商品信息和价格即可提交订单
* 获取用户openid (TODO)

--

调用方式

```GO
package main

import (
	"github.com/erienniu/Ewepay"
)

//type UnifyOrderReq struct {
//	Appid            string `xml:"appid"`
//	Body             string `xml:"body"`
//	Mch_id           string `xml:"mch_id"`
//	Nonce_str        string `xml:"nonce_str"`
//	Notify_url       string `xml:"notify_url"`
//	Trade_type       string `xml:"trade_type"`
//	Spbill_create_ip string `xml:"spbill_create_ip"`
//	Total_fee        int    `xml:"total_fee"`
//	Out_trade_no     string `xml:"out_trade_no"`
//	Sign             string `xml:"sign"`
//}

//载入配置和模式
func main(){
	Ewepay.Config(".properties", "DEVELOPMENT")
	
	/*
	  准备所需的XML字符串和结构体
	  参数分别为商品名称，客户openid，商品价格（单位为分）
	  stru为结构体，可以进行数据库操作
	*/
	str, stru := Ewepay.PrepareSubscriptionXML("商品","wxxxxxxxx")
	
	//提交订单，参数为XML字符串
	Ewepay.SubscriptionOrder(str)
	
}



```