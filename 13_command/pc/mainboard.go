package pc

type MainBoardApi interface {
	// 开机
	Open()

	// 重启
	Reset()
}
