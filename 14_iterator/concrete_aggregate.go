package iterator

type ConcreteAggregate struct {
	ss []string
}

func (c *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(c)
}

func (c *ConcreteAggregate) size() int {
	return len(c.ss)
}

func (c *ConcreteAggregate) get(i int) string {
	if i < c.size() {
		return c.ss[i]
	}
	return ""
}
