package pc

type MotherBoard struct {
	cdDriver  *CDDriver
	cpu       *CPU
	videoCard *VideoCard
	soundCard *SoundCard
}

func (m *MotherBoard) SetCDDriver(c *CDDriver) *MotherBoard {
	m.cdDriver = c
	return m
}

func (m *MotherBoard) SetCPU(c *CPU) *MotherBoard {
	m.cpu = c
	return m
}

func (m *MotherBoard) SetVideoCard(v *VideoCard) *MotherBoard {
	m.videoCard = v
	return m
}

func (m *MotherBoard) SetSoundCard(s *SoundCard) *MotherBoard {
	m.soundCard = s
	return m
}
func (m *MotherBoard) Change(c Colleague) {
	switch c.(type) {
	case *CDDriver:
		m.opeCDDriverReadData(c.(*CDDriver))
	case *CPU:
		m.opeCPU(c.(*CPU))
	}
}

func (m *MotherBoard) opeCDDriverReadData(cd *CDDriver) {
	data := cd.GetData()
	m.cpu.ExecuteData(data)
}

func (m *MotherBoard) opeCPU(cpu *CPU) {
	// cpu处理后的音频视频数据
	soundData := m.cpu.GetSoundData()
	videoData := m.cpu.GetVideoData()
	// 播放
	m.videoCard.ShowData(videoData)
	m.soundCard.SoundData(soundData)
}
