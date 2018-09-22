package alipay

import (
	"errors"
	"net/url"

	"github.com/samhuangszu/alipay/encoding"
)

// VerifyGateWay 网关验证
func (c *AliPay) VerifyGateWay(values url.Values) (string, error) {
	signType := values.Get("sign_type")
	if len(signType) <= 0 {
		return "", errors.New("signType 为空")
	}
	bizContent := values.Get("biz_content")
	if len(bizContent) <= 0 {
		return "", errors.New("biz_content 为空")
	}
	charset := values.Get("charset")
	ok, err := c.VerifySign(values, true)
	if err != nil {
		return "", err
	}
	if !ok {
		return c.VerifyRespXML(bizContent, signType, charset, false)
	}
	return c.VerifyRespXML(bizContent, signType, charset, true)
}

// VerifyRespXML 返回xml字符串
func (c *AliPay) VerifyRespXML(bizContent string, signType string, charset string, result bool) (string, error) {
	// 组装XML
	xml := ""
	publicKey := encoding.OrginPublicKey(string(c.appPublicKey))
	if result == true {
		xml = "<success>true</success><biz_content>" + publicKey + "</biz_content>"
	} else {
		xml = "<success>false</success><error_code>VERIFY_FAILED</error_code><biz_content>" + publicKey + "</biz_content>"
	}
	sign, err := signWithString(xml, c.appPrivateKey, signType)
	if err != nil {
		return "", err
	}
	respXML := "<?xml version=\"1.0\" encoding=\"" + charset + "\"?><alipay><response>" + xml + "</response><sign>" + sign + "</sign><sign_type>" + signType + "</sign_type></alipay>"
	return respXML, nil
}
