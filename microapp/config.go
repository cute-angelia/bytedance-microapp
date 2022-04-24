package microapp

// config options
type config struct {
	AppId     string
	AppSecret string
}

// DefaultConfig 返回默认配置
func DefaultConfig() *config {
	return &config{}
}
