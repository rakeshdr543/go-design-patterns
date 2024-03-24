package decorator

type VegPizza struct {
}

func (p *VegPizza) GetPrice() float64 {
	return 7.5
}
