package pc

import "testing"

func TestRun(t *testing.T) {
	// 创建中介者 motherboard
	motherBoard := &MotherBoard{}

	// 创建同事类
	cdDriver := NewCDDriver(motherBoard)
	cpu := NewCPU(motherBoard)
	videoCard := NewVideoCard(motherBoard)
	soundCard := NewSoundCard(motherBoard)

	// 让中介者知道所有的同事
	motherBoard.SetCDDriver(cdDriver).SetCPU(cpu).
		SetSoundCard(soundCard).
		SetVideoCard(videoCard)

	// 开始读取CD内容
	cdDriver.ReadCD()
}
