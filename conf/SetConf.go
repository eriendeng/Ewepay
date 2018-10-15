package conf

import "errors"

var config = Config{}
var isSet = false

type Config struct {
	Version string
	App int
	AppId string
	MchId string
	Secret string
	NotifyUrl string
	WxPayApiKey string
}

func (c C) Use(node string) {
	config.Version = c.Read(node, "VERSION")
	switch (c.Read(node, "APP")){
	case "subscription":
		config.App = 0
	case "miniprogram":
		config.App = 1
	default:
		config.App = 0
	}
	config.AppId = c.Read(node, "APPID")
	config.MchId = c.Read(node, "MCHID")
	config.Secret = c.Read(node, "SECRET")
	config.NotifyUrl = c.Read(node, "NOTIFYURL")
	config.WxPayApiKey = c.Read(node, "WXPAYAPIKEY")
	isSet = true
}

func GetConfig() (Config){
	if !isSet {
		panic(errors.New("Configruation is not INIT"))
	}
	return config
}