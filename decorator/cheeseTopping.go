package decorator

type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() float64 {
	return c.Pizza.GetPrice() + 2.5
}
