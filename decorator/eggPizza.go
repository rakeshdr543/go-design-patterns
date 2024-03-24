package decorator

type EggPizza struct {
}

func (p *EggPizza) GetPrice() float64 {
	return 9.5
}
