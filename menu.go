package alipay

import "errors"

// Menu 设置菜单
func (c *AliPay) Menu(menu *Menu) (*MenuResp, error) {
	menuResp := &MenuResp{}
	err := c.doRequest("POST", menu, menuResp)
	if err != nil {
		return nil, err
	}
	if menuResp.IsSuccess() == false {
		return nil, errors.New(menuResp.Body.Msg)
	}
	return menuResp, nil
}

// QueryMenu 查询菜单
func (c *AliPay) QueryMenu(queryMenu *QueryMenuReq) (*QueryMenuResp, error) {
	queryMenuResp := &QueryMenuResp{}
	err := c.doRequest("POST", queryMenu, queryMenuResp)
	if err != nil {
		return nil, err
	}
	if queryMenuResp.IsSuccess() == false {
		return nil, errors.New(queryMenuResp.Body.Msg)
	}
	return queryMenuResp, nil
}
