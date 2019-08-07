package composite

import (
	"testing"
)

func TestRun(t *testing.T) {
	root := NewComposite("服装")
	c1 := NewComposite("男装")
	c2 := NewComposite("女装")
	leaf1 := NewLeaf("衬衣")
	leaf2 := NewLeaf("夹克")
	leaf3 := NewLeaf("裙子")
	leaf4 := NewLeaf("套装")
	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(leaf1)
	c1.AddChild(leaf2)
	c2.AddChild(leaf3)
	c2.AddChild(leaf4)

	root.PrintStruct("")
}
