package pc

import "strings"

type CPU struct {
	mediator  Mediator
	videoData string
	soundData string
}

func NewCPU(m Mediator) *CPU {
	return &CPU{
		mediator: m,
	}
}
func (c *CPU) GetMediator() Mediator {
	return c.mediator
}
func (c *CPU) GetVideoData() string {
	return c.videoData
}

func (c *CPU) GetSoundData() string {
	return c.soundData
}
func (c *CPU) ExecuteData(data string) {
	ss := strings.SplitN(data, ",", -1)
	c.videoData = ss[0]
	c.soundData = ss[1]
	c.GetMediator().Change(c)
}
