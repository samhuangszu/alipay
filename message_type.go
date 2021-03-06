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
	HeadColor  string   `json:"head_color"`  //顶部色条的色值
	URL        string   `json:"url"`         // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	ActionName string   `json:"action_name"` //底部链接描述文字，如“查看详情”
	Keyword1   DataItem `json:"keyword1,omitempty"`
	Keyword2   DataItem `json:"keyword2,omitempty"`
	Keyword3   DataItem `json:"keyword3,omitempty"`
	Keyword4   DataItem `json:"keyword4,omitempty"`
	Keyword5   DataItem `json:"keyword5,omitempty"`
	Keyword6   DataItem `json:"keyword6,omitempty"`
	Keyword7   DataItem `json:"keyword7,omitempty"`
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

// IndustryReq 设置行业
// https://docs.open.alipay.com/api_6/alipay.open.public.template.message.industry.modify
type IndustryReq struct {
	AppAuthToken          string `json:"-"` //授权凭证
	PrimaryIndustryName   string `json:"primary_industry_name"`
	PrimaryIndustryCode   string `json:"primary_industry_code"`
	SecondaryIndustryCode string `json:"secondary_industry_code"`
	SecondaryIndustryName string `json:"secondary_industry_name"`
}

// APIName 接口名字
func (c IndustryReq) APIName() string {
	return "alipay.open.public.template.message.industry.modify"
}

// Params 请求参数
func (c IndustryReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c IndustryReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c IndustryReq) ExtJSONParamValue() string {
	return marshal(c)
}

// IndustryResp 响应
type IndustryResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_open_public_template_message_industry_modify_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否成功
func (c IndustryResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

//IndustryQueryReq 查看是否设置行业
// https://docs.open.alipay.com/api_6/alipay.open.public.setting.category.query/
type IndustryQueryReq struct {
	AppAuthToken string `json:"-"` //授权凭证
}

// APIName 接口名字
func (c IndustryQueryReq) APIName() string {
	return "alipay.open.public.setting.category.query"
}

// Params 请求参数
func (c IndustryQueryReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c IndustryQueryReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c IndustryQueryReq) ExtJSONParamValue() string {
	return marshal(c)
}

// IndustryQueryResp 响应
type IndustryQueryResp struct {
	Body struct {
		Code              string `json:"code"`
		Msg               string `json:"msg"`
		SubCode           string `json:"sub_code"`
		SubMsg            string `json:"sub_msg"`
		PrimaryCategory   string `json:"primary_category"`
		SecondaryCategory string `json:"secondary_category"`
	} `json:"alipay_open_public_setting_category_query_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否成功
func (c IndustryQueryResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE && len(c.Body.PrimaryCategory) > 0 {
		return true
	}
	return false
}

// GroupMessageReq 群发消息
// https://docs.open.alipay.com/api_6/alipay.open.public.message.total.send/
type GroupMessageReq struct {
	AppAuthToken string     `json:"-"`        //授权凭证
	MsgType      string     `json:"msg_type"` //消息类型，text：文本消息，image-text：图文消息
	Articles     []*Article `json:"articles,omitempty"`
	Text         *Text      `json:"text,omitempty"`
}

// Article 图文
type Article struct {
	Title      string `json:"title,omitemtpy"`
	Desc       string `json:"desc"`
	ImageURL   string `json:"image_url"`
	URL        string `json:"url"`
	ActionName string `json:"action_name,omitempty"` //链接文字
}

// Text 文本
type Text struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// APIName 接口名字
func (c GroupMessageReq) APIName() string {
	return "alipay.open.public.message.total.send"
}

// Params 请求参数
func (c GroupMessageReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// ExtJSONParamName ext字段名字
func (c GroupMessageReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue ext字段内容
func (c GroupMessageReq) ExtJSONParamValue() string {
	return marshal(c)
}

// GroupMessageResp 响应
type GroupMessageResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_open_public_message_total_send_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否成功
func (c GroupMessageResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
