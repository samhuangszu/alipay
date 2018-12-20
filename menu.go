package alipay

import "errors"

// Menu 设置菜单
func (c *AliPay) Menu(menu *Menu) (MenuResp, error) {
	var resp MenuResp
	if menu.IsCreated {
		resp = &MenuCreateResp{}
	} else {
		resp = &MenuModifyResp{}
	}
	err := c.doRequest("POST", menu, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() == false {
		return nil, errors.New(resp.Msg())
	}
	return resp, nil
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
