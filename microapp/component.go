package microapp

type Component struct {
	config *config
}

func newComponent(config *config) *Component {
	comp := &Component{}
	comp.config = config
	return comp
}

// GetAccessToken
// https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/interface-request-credential/get-access-token
func (c *Component) GetAccessToken() {

}
