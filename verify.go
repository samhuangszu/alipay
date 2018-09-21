package alipay

import (
	"encoding/xml"
	"errors"
	"net/url"
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
	ok, err := c.VerifySign(values)
	if err != nil {
		return "", err
	}
	if !ok {
		return c.VerifyRespXML(bizContent, signType, false)
	}
	return c.VerifyRespXML(bizContent, signType, true)
}

// VerifyRespXML 返回xml字符串
func (c *AliPay) VerifyRespXML(bizContent string, signType string, result bool) (string, error) {
	aliBizContent := &AliBizContent{}
	err := xml.Unmarshal([]byte(bizContent), aliBizContent)
	if err != nil {
		return "", err
	}
	// 组装XML
	xml := ""
	if result == true {
		xml = "<success>true</success><biz_content>" + string(c.AliPayPublicKey) + "</biz_content>"
	} else {
		xml = "<success>false</success><error_code>VERIFY_FAILED</error_code><biz_content>" + string(c.AliPayPublicKey) + "</biz_content>"
	}
	sign, err := signWithString(xml, c.AliPayPublicKey, signType)
	if err != nil {
		return "", err
	}
	respXML := "<?xml version=\"1.0\" encoding=\"" + K_CHARSET + "\"?><alipay><response>" + xml + "</response><sign>" + sign + "</sign><sign_type>" + signType + "</sign_type></alipay>"
	return respXML, nil
}
