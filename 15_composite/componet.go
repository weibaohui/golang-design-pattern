package composite

type Component interface {
	PrintStruct(pre string)
	AddChild(child Component)
	RemoveChild(child Component)
	GetChild(index int) Component
}
