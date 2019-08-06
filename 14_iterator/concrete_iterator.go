package iterator

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int // 内部索引
}

func (c *ConcreteIterator) First() {
	c.index = 0
}

func (c *ConcreteIterator) Next() {
	if c.index < c.aggregate.size() {
		c.index += 1
	}
}

func (c *ConcreteIterator) IsDone() bool {
	return c.index == c.aggregate.size()
}

func (c *ConcreteIterator) CurrentItem() interface{} {
	return c.aggregate.get(c.index)
}

func NewConcreteIterator(aggregate *ConcreteAggregate) Iterator {
	return &ConcreteIterator{aggregate: aggregate}
}
