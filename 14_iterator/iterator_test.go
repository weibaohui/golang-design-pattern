package iterator

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	names := []string{"张三", "李四", "王五"}
	aggregate := NewConcreteAggregate(names)
	iterator := aggregate.CreateIterator()

	iterator.First()
	for !iterator.IsDone() {
		item := iterator.CurrentItem()
		fmt.Println(item)
		iterator.Next()
	}
}
