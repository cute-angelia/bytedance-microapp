package apis

import (
	"github.com/cute-angelia/bytedance-microapp/microapp/utils"
	"github.com/guonaihong/gout"
)

// SendSubscribeNotification 发送订阅消息
// 用户产生了订阅模板消息的行为后，可以通过这个接口发送模板消息给用户。
// see https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/subscribe-notification/notify
func SendSubscribeNotification(accessToken, appId, tplId, openid string, data map[string]string, page string) (resp SendSubscribeNotificationResp) {
	utils.HttpClient.POST("https://developer.toutiao.com/api/apps/subscribe_notification/developer/v1/notify").SetJSON(gout.H{
		"access_token": accessToken, //小程序 access_token
		"app_id":       appId,       //小程序的 id
		"tpl_id":       tplId,       //模板的 id
		"open_id":      openid,      //接收消息目标用户的 open_id
		"data":         data,        //模板内容，格式形如 { "key1": "value1", "key2": "value2" }
		"page":         page,        //跳转的页面, 可为空
	}).BindJSON(&resp).Do()
	return resp
}

type SendSubscribeNotificationResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
}
