package composite

import "fmt"

type Composite struct {
	Name            string
	childComponents []Component
}

func NewComposite(name string) *Composite {
	return &Composite{
		Name:            name,
		childComponents: make([]Component, 0),
	}
}
func (c *Composite) PrintStruct(pre string) {
	fmt.Println(pre + " + " + c.Name)
	pre += "  "
	for _, cc := range c.childComponents {
		cc.PrintStruct(pre)
	}
}

func (c *Composite) AddChild(child Component) {
	c.childComponents = append(c.childComponents, child)
}

func (c *Composite) RemoveChild(child Component) {
	for i := 0; i < len(c.childComponents); i++ {
		if c.childComponents[i] == child {
			c.childComponents = append(c.childComponents[:i], c.childComponents[i+1:]...)
			i--
		}
	}
}

func (c *Composite) GetChild(index int) Component {
	if index <= len(c.childComponents) {
		return c.childComponents[index]
	}
	return nil
}
