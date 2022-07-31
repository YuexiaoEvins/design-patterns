package Decorator

type IPizza interface {
	GetPrice() int
}

type PizzaHub struct {
}

func (p *PizzaHub) GetPrice() int {
	return 10
}

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	p := t.pizza.GetPrice()
	return p + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	p := c.pizza.GetPrice()
	return p + 15
}
