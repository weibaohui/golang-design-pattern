package composite

import "fmt"

type Leaf struct {
	Name string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{Name: name}
}
func (c *Leaf) PrintStruct(pre string) {
	fmt.Println(pre + " - " + c.Name)
}

func (c *Leaf) AddChild(child Component) {
	fmt.Println("无 AddChild 功能")
}

func (c *Leaf) RemoveChild(child Component) {
	fmt.Println("无 RemoveChild 功能")
}

func (c *Leaf) GetChild(index int) Component {
	fmt.Println("无 GetChild 功能")
	return nil
}
