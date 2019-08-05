package pc

import "testing"

func TestRun(t *testing.T) {
	mainBoard := &GigaMainBoard{}

	openCommand := NewOpenCommand(mainBoard)
	resetCommand := NewResetCommand(mainBoard)

	// 让按钮知道该干什么
	box := &Box{}
	box.SetOpenCommand(openCommand)
	box.SetResetCommand(resetCommand)

	// 模拟按按钮
	box.OpenPressed()
	box.ResetPressed()
}
