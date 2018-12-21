package alipay

import "errors"

// SendMessage 发消息
func (c *AliPay) SendMessage(req *SendMessageReq) (resp *SendMessageResp, err error) {
	err = c.doRequest("POST", req, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() == false {
		return nil, errors.New(resp.Body.Msg)
	}
	return resp, nil
}

// GetTemplate 领取模版
func (c *AliPay) GetTemplate(req *GetTemplateReq) (resp *GetTemplateResp, err error) {
	err = c.doRequest("POST", req, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() == false {
		return nil, errors.New(resp.Body.Msg)
	}
	return resp, nil
}
