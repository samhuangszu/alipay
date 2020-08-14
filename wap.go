package alipay

import (
	"net/url"
)

// TradeWapPay 手机网站支付接口 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (c *AliPay) TradeWapPay(param TradeWapPay) (result *url.URL, err error) {
	p, err := c.URLValues(param)
	if err != nil {
		return nil, err
	}
	result, err = url.Parse(c.apiDomain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, err
}
