package alipay

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// TradeWapPay https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (this *AliPay) TradeWapPay(param AliPayTradeWapPay) (resp []byte, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}
	var buf = strings.NewReader(p.Encode())

	req, err := http.NewRequest("POST", this.apiDomain, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", K_CONTENT_TYPE_FORM)

	rep, err := this.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()

	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rep.Body)
	return data, err
}
