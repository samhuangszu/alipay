package alipay

// InfoReq 基本信息
type InfoReq struct {
	AppAuthToken string `json:"-"`
}

// APIName 接口名字
func (c InfoReq) APIName() string {
	return "alipay.open.public.info.query"
}

// Params 请求参数
func (c InfoReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c InfoReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c InfoReq) ExtJSONParamValue() string {
	return marshal(c)
}

// InfoResp 请求响应
type InfoResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
		AppName string `json:"app_name"`
		LogoURL string `json:"logo_url"`
	} `json:"alipay_open_public_info_query_response"`
	Sign string `json:"sign"`
}

// IsSuccess 返回是否成功
func (c *InfoResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
