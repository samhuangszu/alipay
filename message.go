package alipay

// SendMessage 发消息
func (c *AliPay) SendMessage(req *SendMessageReq) (*SendMessageResp, error) {
	resp := &SendMessageResp{}
	err := c.doRequest("POST", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTemplate 领取模版
func (c *AliPay) GetTemplate(req *GetTemplateReq) (*GetTemplateResp, error) {
	resp := &GetTemplateResp{}
	err := c.doRequest("POST", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
