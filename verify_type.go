package alipay

import "encoding/xml"

// AliBizContent 验证网关
type AliBizContent struct {
	XMLName          xml.Name `xml:"XML"`
	AppID            string   `xml:"AppId"`
	FromUserID       string   `xml:"FromUserId"`
	FromAlipayUserID string   `xml:"FromAlipayUserId"`
	CreateTime       string   `xml:"CreateTime"`
	MsgType          string   `xml:"MsgType"`
	EventType        string   `xml:"EventTpye"`
	ActionParam      string   `xml:"ActionParam"`
	AgreementID      string   `xml:"AgreementId"`
	AccountNo        string   `xml:"AccountNo"`
}

// AliForm 阿里返回
type AliForm struct {
	Sign       string `param:"sign"`
	Charset    string `param:"charset"`
	BizContent string `param:"biz_content"`
	SignType   string `param:"sign_type"`
	Service    string `param:"service"`
}

// AliVerifyResp 返回给阿里
type AliVerifyResp struct {
	XMLName  xml.Name `xml:"alipay"`
	Response struct {
		BizContent string `xml:"biz_content"`
		Success    string `xml:"success"`
	} `xml:"response"`
	Sign     string `xml:"sign"`
	SignType string `xml:"sign_type"`
}
