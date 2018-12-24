package alipay

// QRCodeSetReq 设置推广码
// https://docs.open.alipay.com/api_6/alipay.open.public.qrcode.create/
type QRCodeSetReq struct {
	AppAuthToken string `json:"-"`
	CodeInfo     struct {
		Scene struct {
			SceneID string `json:"scene_id,omitempty"`
		} `json:"scene,omitempty"`
		GotoURL string `json:"goto_url,omitempty"`
	} `json:"code_info,omitempty"`
	CodeType     string `json:"code_type,omitempty"`
	ExpireSecond string `json:"expire_second,omitempty"`
	ShowLogo     string `json:"show_logo,omitempty"` // 是否显示logo:Y / N
}

// APIName 接口名字
func (c QRCodeSetReq) APIName() string {
	return "alipay.open.public.qrcode.create"
}

// Params 请求参数
func (c QRCodeSetReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c QRCodeSetReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c QRCodeSetReq) ExtJSONParamValue() string {
	return marshal(c)
}

// QRCodeSetResp 设置推广码
type QRCodeSetResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
		CodeImg string `json:"code_img"`
	} `json:"alipay_open_public_qrcode_create_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否成功
func (c QRCodeSetResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
