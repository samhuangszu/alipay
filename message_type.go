package alipay

//SendMessageReq 发送的模板消息内容
//https://docs.open.alipay.com/api_6/alipay.open.public.message.single.send/
type SendMessageReq struct {
	AppAuthToken string   `json:"-"`          //授权凭证
	ToUserID     string   `json:"to_user_id"` // 必须, ali_user_id
	Template     Template `json:"template"`   //必填，消息内容
}

// Template 消息内容
type Template struct {
	TemplateID string  `json:"template_id"`
	Context    Context `json:"context"`
}

// Context 消息上下文
type Context struct {
	HeadColor  string   `json:"head_color,omitempty"` //顶部色条的色值
	URL        string   `json:"url"`                  // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	ActionName string   `json:"action_name"`          //底部链接描述文字，如“查看详情”
	Keyword1   DataItem `json:"keyword1,omitempty"`
	Keyword2   DataItem `json:"keyword2,omitempty"`
	Keyword3   DataItem `json:"keyword3,omitempty"`
	Keyword4   DataItem `json:"keyword4,omitempty"`
	Keyword5   DataItem `json:"keyword5,omitempty"`
	First      DataItem `json:"first,omitempty"`
	Remark     DataItem `json:"remark,omitempty"`
}

//DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// APIName 接口名字
func (c SendMessageReq) APIName() string {
	return "alipay.open.public.message.single.send"
}

// Params 请求参数
func (c SendMessageReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c SendMessageReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c SendMessageReq) ExtJSONParamValue() string {
	return marshal(c)
}

// SendMessageResp 消息响应
type SendMessageResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_open_public_message_single_send_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否成功
func (c SendMessageResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

// GetTemplateReq 领取模版
// https://docs.open.alipay.com/api_6/alipay.open.public.template.message.get/
type GetTemplateReq struct {
	TemplateID   string `json:"template_id"` //公用模版ID：TM000000223
	AppAuthToken string `json:"-"`           //授权凭证
}

// APIName 接口名字
func (c GetTemplateReq) APIName() string {
	return "alipay.open.public.template.message.get"
}

// Params 请求参数
func (c GetTemplateReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c GetTemplateReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c GetTemplateReq) ExtJSONParamValue() string {
	return marshal(c)
}

// GetTemplateResp 消息响应
type GetTemplateResp struct {
	Body struct {
		Code          string `json:"code"`
		Msg           string `json:"msg"`
		SubCode       string `json:"sub_code"`
		SubMsg        string `json:"sub_msg"`
		Template      string `json:"template"`
		MsgTemplateID string `json:"msg_template_id"`
	} `json:"alipay_open_public_template_message_get_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否成功
func (c GetTemplateResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
