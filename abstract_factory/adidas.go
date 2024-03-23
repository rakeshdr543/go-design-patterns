package abstract_factory

type Adidas struct{}

func (a *Adidas) MakeShoe(size int) IShoe {
	return &Shoe{
		logo: "adidas",
		size: size + 1,
	}
}

func (a *Adidas) MakeShirt(size int) IShirt {
	return &Shirt{
		logo: "adidas",
		size: size + 2,
	}
}
