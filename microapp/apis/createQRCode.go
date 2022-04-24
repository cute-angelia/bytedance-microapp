package apis

import (
	"github.com/cute-angelia/bytedance-microapp/microapp/utils"
	"github.com/guonaihong/gout"
)

const (
	AppNameToutiao     = "toutiao"
	AppNameToutiaoLite = "toutiao_lite"
	AppNameDouyin      = "douyin"
	AppNameDouyinLite  = "douyin_lite"
	AppNamePipixia     = "pipixia"
	AppNameHuoshan     = "huoshan"
	AppNameXigua       = "xigua"
)

// CreateQRCode 获取 AccessToken
// 获取小程序/小游戏的二维码。该二维码可通过任意 app 扫码打开，能跳转到开发者指定的对应字节系 app 内拉起小程序/小游戏， 并传入开发者指定的参数。通过该接口生成的二维码，永久有效，暂无数量限制。
// see https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/qr-code/create-qr-code
func CreateQRCode(accessToken, appname, path string) []byte {
	var resp []byte
	utils.HttpClient.POST("https://developer.toutiao.com/api/apps/qrcode").SetJSON(gout.H{
		"access_token": accessToken,
		"appname":      appname,
		"path":         path,
	}).BindBody(&resp).Do()
	return resp
}
