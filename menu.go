package alipay

// Menu 设置菜单
func (c *AliPay) Menu(menu *Menu) (result *MenuResp, err error) {
	err = c.doRequest("POST", menu, result)
	return result, err
}
