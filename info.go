package alipay

import "errors"

// GetAppInfo 获取App基本信息
func (c *AliPay) GetAppInfo(infoReq *InfoReq) (*InfoResp, error) {
	infoResp := &InfoResp{}
	err := c.doRequest("POST", infoReq, infoResp)
	if err != nil {
		return nil, err
	}
	if infoResp.IsSuccess() == false {
		return nil, errors.New(infoResp.Body.Msg)
	}
	return infoResp, nil
}
