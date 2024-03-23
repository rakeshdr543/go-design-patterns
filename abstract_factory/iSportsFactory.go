package abstract_factory

type ISportsFactory interface {
	MakeShoe(size int) IShoe
	MakeShirt(size int) IShirt
}

func GetSportsFactory(brand string) ISportsFactory {
	if brand == "adidas" {
		return &Adidas{}
	}

	if brand == "nike" {
		return &Nike{}
	}

	return nil
}
