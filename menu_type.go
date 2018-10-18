package alipay

//Button 菜单按钮
type Button struct {
	Name        string    `json:"name,omitempty"`
	ActionType  string    `json:"action_type,omitempty"`
	ActionParam string    `json:"action_param,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	SubButtons  []*Button `json:"sub_button,omitempty"`
}

//Menu https://docs.open.alipay.com/api_6/alipay.open.public.menu.modify
type Menu struct {
	Button []*Button `json:"button"`
}

// APIName 接口名
func (c Menu) APIName() string {
	return "alipay.open.public.menu.modify"
}

// Params 请求参数
func (c Menu) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// ExtJSONParamName biz_content内容
func (c Menu) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue 内容
func (c Menu) ExtJSONParamValue() string {
	return marshal(c)
}

// MenuResp 响应
type MenuResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		MenuKey string `json:"menu_key"`
	} `json:"alipay_open_public_menu_modify_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否请请成功
func (c *MenuResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

// QueryMenuReq 查询菜单
type QueryMenuReq struct {
}

// APIName 接口名
func (c QueryMenuReq) APIName() string {
	return "alipay.open.public.menu.batchquery"
}

// Params 请求参数
func (c QueryMenuReq) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// ExtJSONParamName biz_content内容
func (c QueryMenuReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue 内容
func (c QueryMenuReq) ExtJSONParamValue() string {
	return marshal(c)
}

// QueryMenuResp 查询返回内容
type QueryMenuResp struct {
	Body struct {
		Code  string  `json:"code"`
		Msg   string  `json:"msg"`
		Menus []*Menu `json:"menus"`
	} `json:"alipay_open_public_menu_batchquery_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否请请成功
func (c *QueryMenuResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
