package abstract_factory

type Nike struct{}

func (n *Nike) MakeShoe(size int) IShoe {
	return &Shoe{
		logo: "nike",
		size: size,
	}
}

func (n *Nike) MakeShirt(size int) IShirt {
	return &Shirt{
		logo: "nike",
		size: size,
	}
}
