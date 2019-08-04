package pc

import "fmt"

type SoundCard struct {
	mediator Mediator
	data     string
}

func NewSoundCard(m Mediator) *SoundCard {
	return &SoundCard{
		mediator: m,
	}
}
func (c *SoundCard) SoundData(data string) {
	fmt.Println("画外音：", data)
}
