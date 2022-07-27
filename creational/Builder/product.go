package Builder

type Builder interface {
	buildStepA()
	buildStepB()
	buildStepC()
	reset()
	GetResult() *Product
}

type Product struct {
	fieldA int
	fieldB string
	fieldC uint32
}

type ConcreteBuilder1 struct {
	p *Product
}

func (c *ConcreteBuilder1) reset() {
	c.p = &Product{}
}
func (c *ConcreteBuilder1) buildStepA() {
	c.p.fieldA = 2
}
func (c *ConcreteBuilder1) buildStepB() {
	c.p.fieldB = "test"
}

func (c *ConcreteBuilder1) buildStepC() {
	c.p.fieldC = 3
}

func (c *ConcreteBuilder1) getResult() *Product {
	return c.p
}

type ConcreteBuilder2 struct {
	p *Product
}

func (c *ConcreteBuilder2) reset() {
	c.p = &Product{}
}
func (c *ConcreteBuilder2) buildStepA() {
	c.p.fieldA = 2
}
func (c *ConcreteBuilder2) buildStepB() {
	c.p.fieldB = "another test"
}

func (c *ConcreteBuilder2) buildStepC() {
	c.p.fieldC = 3
}

func (c *ConcreteBuilder2) getResult() *Product {
	return c.p
}
