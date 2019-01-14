package alipay

// ModifyIndustry 更改行业
func (c *AliPay) ModifyIndustry(req *IndustryReq) (*IndustryResp, error) {
	resp := &IndustryResp{}
	if err := c.doRequest("POST", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryIndustry 查询行业
func (c *AliPay) QueryIndustry(req *IndustryQueryReq) (*IndustryQueryResp, error) {
	resp := &IndustryQueryResp{}
	if err := c.doRequest("POST", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

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
