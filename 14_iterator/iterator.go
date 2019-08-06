package iterator

type Iterator interface {
	// 移动到第一个位置
	First()
	// 移动到下一个位置
	Next()
	// 判断是否已经是最后一个位置
	IsDone() bool
	CurrentItem() interface{}
}
