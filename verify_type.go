package alipay

import "encoding/xml"

// AliVerify 验证网关
type AliVerify struct {
	XMLName     xml.Name `xml:"XML"`
	AppID       string   `xml:"AppId"`
	FromUserID  string   `xml:"FromUserId"`
	CreateTime  string   `xml:"CreateTime"`
	MsgType     string   `xml:"MsgType"`
	EventType   string   `xml:"EventTpye"`
	ActionParam string   `xml:"ActionParam"`
	AgreementID string   `xml:"AgreementId"`
	AccountNo   string   `xml:"AccountNo"`
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
