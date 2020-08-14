package alipay

import "encoding/json"

// TradeWapPay h5支付
// https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
type TradeWapPay struct {
	TradePay
	QuitURL    string `json:"quit_url,omitempty"`
	AuthToken  string `json:"auth_token,omitempty"`  // 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	TimeExpire string `json:"time_expire,omitempty"` // 绝对超时时间，格式为yyyy-MM-dd HH:mm。
}

// APIName apiname
func (c TradeWapPay) APIName() string {
	return "alipay.trade.wap.pay"
}

// Params 参数
func (c TradeWapPay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AuthToken
	m["notify_url"] = c.NotifyURL
	m["return_url"] = c.ReturnURL
	return m
}

// ExtJSONParamName json
func (c TradeWapPay) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue biz_content内容
func (c TradeWapPay) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}
