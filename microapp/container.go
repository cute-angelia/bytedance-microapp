package microapp

import (
	"github.com/cute-angelia/go-utils/syntax/ijson"
	"github.com/spf13/viper"
	"log"
)

type Option func(c *Container)

// Container 两种方式
// 一种从 viper 获取配置
// 一种是 options 模式
type Container struct {
	config *config
}

// Load viper 加载 配置
func Load(key string) *Component {
	iconfig := DefaultConfig()
	configData := viper.GetStringMap(key)
	jsonstr, _ := ijson.Marshal(configData)
	if err := ijson.Unmarshal(jsonstr, &iconfig); err != nil {
		log.Println(err)
	}
	// log.Println(ijson.Pretty(iconfig))
	return newComponent(iconfig)
}

// New options 模式
func New(options ...Option) *Component {
	c := &Container{
		config: DefaultConfig(),
	}
	for _, option := range options {
		option(c)
	}
	return newComponent(c.config)
}

func WithAppId(appId string) Option {
	return func(c *Container) {
		c.config.AppId = appId
	}
}

func WithAppSecret(appSecret string) Option {
	return func(c *Container) {
		c.config.AppSecret = appSecret
	}
}
