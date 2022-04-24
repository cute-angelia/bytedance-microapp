# cute-angelia/bytedance-microapp

字节跳动小程序，抖音小程序

### 快速开始 & demo

```shell
go get github.com/cute-angelia/bytedance-microapp
```

```go
// 初始化
// 方法一 
microapp.New(microapp.WithAppId(""), microapp.WithAppSecret(""))

// 方法二 采用配置文件
conf.LoadConfigFile("./config.toml")
microapp.Load("microapp")

//  调用 Api 如 GetAccessToken
microApp := microapp.New(microapp.WithAppId(""), microapp.WithAppSecret(""))
accessToken := microApp.GetAccessToken()
```

### 登录流程

1. 不授权登陆

   不授权登陆只需要客户端获取 code
    1. 客户端获取 tt.login 获取 code
    2. 服务端调用 code2session 根据 openid 记录用户

2. 授权登陆

   需要用户点击授权
   1. 客户端获取 tt.login 获取 code
   2. 客户端调用 tt.getUserInfo 获取 encryptedData 和 iv 
   3. 将这些数据发送服务端进行解密


