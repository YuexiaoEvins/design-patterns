package Builder

//TODO 需要再深化builder模型
type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) registryBuilder(b Builder) {
	d.builder = b
}

func (d *Director) BuildProduct() *Product {
	d.builder.buildStepA()
	d.builder.buildStepB()
	d.builder.buildStepC()
	return d.builder.GetResult()
}
