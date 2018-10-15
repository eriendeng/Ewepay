package order

import (
	"fmt"
	"sort"
	"crypto/md5"
	"strings"
	"encoding/hex"
	"../conf"
	"encoding/xml"
	"bytes"
	"net/http"
	"io/ioutil"
)

func SubscriptionOrder(str string)  {
	bytes_req := []byte(str)
	req, err := http.NewRequest("POST", OrderPostUrl, bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return
	}
	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, _err := c.Do(req)
	if _err != nil {
		fmt.Println("请求微信支付统一下单接口发送错误, 原因:", _err)
		return
	}
	xmlResp := UnifyOrderResp{}
	resp_body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string((resp_body)))
	_err = xml.Unmarshal(resp_body, &xmlResp)
	//处理return code.
	if xmlResp.Return_code == "FAIL" {
		fmt.Println("微信支付统一下单不成功，原因:", xmlResp.Return_msg)
		return
	}

	//这里已经得到微信支付的prepay id，需要返给客户端，由客户端继续完成支付流程
	fmt.Println("微信支付统一下单成功，预支付单号:", xmlResp.Prepay_id)
}


func PrepareSubscriptionXML(name, openid string, fee int) (str string, req UnifyOrderReq){
	var Req UnifyOrderReq
	Req.Appid = conf.GetConfig().AppId
	Req.Body = name//商品名
	Req.Mch_id = conf.GetConfig().MchId
	Req.Nonce_str = RandomString()
	Req.Notify_url = conf.GetConfig().NotifyUrl
	Req.Trade_type = "JSAPI"
	Req.Spbill_create_ip = "xxx.xxx.xxx.xxx"
	Req.Total_fee = fee //单位是分
	Req.Out_trade_no = OrderNumber()

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = Req.Appid
	m["body"] = Req.Body
	m["mch_id"] = Req.Mch_id
	m["notify_url"] = Req.Notify_url
	m["openid"] = openid
	m["trade_type"] = Req.Trade_type
	m["spbill_create_ip"] = Req.Spbill_create_ip
	m["total_fee"] = Req.Total_fee
	m["out_trade_no"] = Req.Out_trade_no
	m["nonce_str"] = Req.Nonce_str
	Req.Sign = WxpayCalcSign(m, conf.GetConfig().WxPayApiKey)

	bytes_req, _ := xml.Marshal(Req)
	//if err != nil {
	//	fmt.Println("以xml形式编码发送错误, 原因:", err)
	//	return
	//}

	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)

	//返回可用于直接post的字符串和用于存放数据库的结构体
	return str_req, Req
}

//微信支付计算签名的函数
func WxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

