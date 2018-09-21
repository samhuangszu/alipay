package alipay

import "encoding/json"

// UserReq 获取access_token及UserID
type UserReq struct {
	GrantType    string `json:"grant_type"` //authorization_code或refresh_token
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// APIName 接口名
func (c UserReq) APIName() string {
	return "alipay.system.oauth.token"
}

// Params 请求参数
func (c UserReq) Params() map[string]string {
	var m = make(map[string]string)
	m["grant_type"] = c.GrantType
	m["code"] = c.Code
	m["refresh_token"] = c.RefreshToken
	return m
}

// ExtJSONParamName 返回名称
func (c UserReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue 返回内容
func (c UserReq) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// UserResp 返回值
type UserResp struct {
	Body struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		AccessToken  string `json:"access_token"`
		UserID       string `json:"user_id"`
		AlipayUserID string `json:"alipay_user_id"` //用户的open_id（已废弃，请勿使用）
		ExpiresIn    string `json:"expires_in"`     //令牌有效期：300秒
		ReExpiresIn  string `json:"re_expires_in"`  // 刷新令牌有效期：300秒
		RefreshToken string `json:"refresh_token"`  // 刷新令牌
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"`
}

// IsSuccess 获取是否成功
func (c *UserResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

// UserInfoReq 用户详情请求
type UserInfoReq struct {
	AuthToken string `json:"auth_token"` //对应到access_token
}

// APIName 接口名
func (c UserInfoReq) APIName() string {
	return "alipay.user.info.share"
}

// Params 请求参数
func (c UserInfoReq) Params() map[string]string {
	var m = make(map[string]string)
	m["auth_token"] = c.AuthToken
	return m
}

// ExtJSONParamName 返回名称
func (c UserInfoReq) ExtJSONParamName() string {
	return "biz_content"
}

// ExtJSONParamValue 返回内容
func (c UserInfoReq) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// UserInfoResp 返回
type UserInfoResp struct {
	Body struct {
		Code               string `json:"code"`
		Msg                string `json:"msg"`
		UserID             string `json:"user_id"`
		Avatar             string `json:"avatar"`
		Province           string `json:"province"`
		City               string `json:"city"`
		IsCertified        string `json:"is_certified"`
		Nickname           string `json:"nick_name"`
		UserStatus         string `json:"user_status"`
		UserType           string `json:"user_type"`
		IsStudentCertified string `json:"is_student_certified"`
		Gender             string `json:"gender"` // 性别
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"`
}

// IsSuccess 获取是否成功
func (c *UserInfoResp) IsSuccess() bool {
	if c.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
