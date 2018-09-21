package alipay

// User 获取user_id,access_token;method:alipay.system.oauth.token
func (c *AliPay) User(param UserReq) (results *UserResp, err error) {
	err = c.doRequest("POST", param, &results)
	return results, err
}

// UserInfo 获取用户详细信息
func (c *AliPay) UserInfo(param UserInfoReq) (results *UserInfoResp, err error) {
	err = c.doRequest("POST", param, &results)
	return results, err
}
