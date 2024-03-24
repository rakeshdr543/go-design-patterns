package decorator

type MushroomTopping struct {
	Pizza IPizza
}

func (m *MushroomTopping) GetPrice() float64 {
	return m.Pizza.GetPrice() + 3.5
}
