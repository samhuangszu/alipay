package alipay

import "fmt"

//Button 菜单按钮
type Button struct {
	Name        string    `json:"name,omitempty"`
	ActionType  string    `json:"action_type,omitempty"`
	ActionParam string    `json:"action_param,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	SubButtons  []*Button `json:"sub_button,omitempty"`
}

//Menu https://docs.open.alipay.com/api_6/alipay.open.public.menu.create
type Menu struct {
	Button []*Button `json:"button"`
}

// APIName 接口名
func (c Menu) APIName() string {
	return "alipay.open.public.menu.create"
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
	fmt.Printf("alipay.Menu.ExtJSONParamValue:%s", marshal(c))
	return marshal(c)
}

// MenuResp 响应
type MenuResp struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		MenuKey string `json:"menu_key"`
	} `json:"alipay_open_public_menu_create_response"`
	Sign string `json:"sign"`
}

// IsSuccess 是否请请成功
func (c *MenuResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
