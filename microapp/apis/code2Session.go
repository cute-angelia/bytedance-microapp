package apis

import (
	"github.com/cute-angelia/bytedance-microapp/microapp/utils"
	"github.com/guonaihong/gout"
)

// Code2Session 获取 AccessToken
// 通过login接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId。
// 登录凭证 code，anonymous_code 只能使用一次，非匿名需要 code，非匿名下的 anonymous_code 用于数据同步，匿名需要 anonymous_code。
// see https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session
func Code2Session(appid, secret string, code, anonymousCode string) Code2SessionResp {
	var resp Code2SessionResp
	utils.HttpClient.POST("https://developer.toutiao.com/api/apps/v2/jscode2session").SetJSON(gout.H{
		"appid":          appid,
		"secret":         secret,
		"code":           code,
		"anonymous_code": anonymousCode,
	}).BindJSON(&resp).Do()
	return resp
}

type Code2SessionResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		SessionKey      string `json:"session_key"`
		Openid          string `json:"openid"`
		AnonymousOpenid string `json:"anonymous_openid"`
		Unionid         string `json:"unionid"`
	} `json:"data"`
}
