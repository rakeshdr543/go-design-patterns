package main

// factory "github.com/rakeshdr543/go-design-patterns/factory"
// af "github.com/rakeshdr543/go-design-patterns/abstract_factory"
// "github.com/rakeshdr543/go-design-patterns/singleton"
import (
	observer "github.com/rakeshdr543/go-design-patterns/observer"
)

func main() {
	// Factory Pattern
	// ak47, _ := factory.GetGun("ak47")
	// musket, _ := factory.GetGun("musket")

	// printDetails(ak47)
	// printDetails(musket)

	// Abstract Factory Pattern
	// nikeFactory := af.GetSportsFactory("nike")
	// adidasFactory := af.GetSportsFactory("adidas")

	// nikeShoe := nikeFactory.MakeShoe(10)
	// nikeShirt := nikeFactory.MakeShirt(20)

	// adidasShoe := adidasFactory.MakeShoe(10)
	// adidasShirt := adidasFactory.MakeShirt(20)

	// printShoeDetails(nikeShoe)
	// printShirtDetails(nikeShirt)

	// printShoeDetails(adidasShoe)
	// println()
	// printShirtDetails(adidasShirt)

	// Singleton pattern

	// for i := 0; i < 10; i++ {
	// 	go singleton.GetInstance()
	// }

	// fmt.Scanln()

	// Observer pattern
	appleMonitor := observer.NewItem("Apple Monitor")
	customer1 := observer.NewCustomer("rakesh1@gmail.com")
	customer2 := observer.NewCustomer("ramesh2@gmail.com")

	appleMonitor.Register(customer1)
	appleMonitor.Register(customer2)

	appleMonitor.UpdateAvailability()

}

// func printShoeDetails(shoe af.IShoe) {
// 	println("Logo: ", shoe.GetLogo())
// 	println("Size: ", shoe.GetSize())
// 	println()
// }

// func printShirtDetails(shirt af.IShirt) {
// 	println("Logo: ", shirt.GetLogo())
// 	println("Size: ", shirt.GetSize())
// 	println()

// }
