package apis

import (
	"github.com/cute-angelia/bytedance-microapp/microapp/utils"
	"github.com/guonaihong/gout"
)

// GetAccessToken 获取 AccessToken
// access_token 是小程序的全局唯一调用凭据，开发者调用小程序支付时需要使用 access_token。
// access_token 的有效期为 2 个小时，需要定时刷新 access_token，重复获取会导致之前一次获取的 access_token 的有效期缩短为 5 分钟。
// see https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/interface-request-credential/get-access-token
func GetAccessToken(appid, secret string) GetAccessTokenResp {
	var resp GetAccessTokenResp
	utils.HttpClient.POST("https://developer.toutiao.com/api/apps/v2/token").SetJSON(gout.H{
		"appid":      appid,
		"secret":     secret,
		"grant_type": "client_credential",
	}).BindJSON(&resp).Do()
	return resp
}

type GetAccessTokenResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
}
