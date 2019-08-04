package pc

type CDDriver struct {
	mediator Mediator
	data     string
}

func NewCDDriver(m Mediator) *CDDriver {
	return &CDDriver{
		mediator: m,
	}
}
func (c *CDDriver) GetMediator() Mediator {
	return c.mediator
}
func (c *CDDriver) GetData() string {
	return c.data
}
func (c *CDDriver) ReadCD() {
	c.data = "视频数据,音频数据"
	c.GetMediator().Change(c)
}
