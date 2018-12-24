package alipay

// QRCodeSet 设置推广码
func (c *AliPay) QRCodeSet(req *QRCodeSetReq) (*QRCodeSetResp, error) {
	resp := &QRCodeSetResp{}
	err := c.doRequest("POST", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
