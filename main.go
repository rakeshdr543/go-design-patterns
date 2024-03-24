package main

import (
	adapter "github.com/rakeshdr543/go-design-patterns/adapter"
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
	// appleMonitor := observer.NewItem("Apple Monitor")
	// customer1 := observer.NewCustomer("rakesh1@gmail.com")
	// customer2 := observer.NewCustomer("ramesh2@gmail.com")

	// appleMonitor.Register(customer1)
	// appleMonitor.Register(customer2)

	// appleMonitor.UpdateAvailability()

	// Builder pattern
	// normalBuilder := builder.GetHouseBuilder("normal")
	// iglooBuilder := builder.GetHouseBuilder("igloo")

	// director := builder.NewDirector(normalBuilder)
	// normalHouse := director.BuildHouse()

	// director.SetBuilder(iglooBuilder)
	// iglooHouse := director.BuildHouse()

	// fmt.Println("Normal House  :", normalHouse)
	// fmt.Println("Igloo House  :", iglooHouse)

	// Decorator pattern
	// vegPizza := &decorator.VegPizza{}
	// cheesePizza := &decorator.CheeseTopping{
	// 	Pizza: vegPizza,
	// }

	// mushroomPizza := &decorator.MushroomTopping{
	// 	Pizza: cheesePizza,
	// }

	// println("Price of veg pizza: ", vegPizza.GetPrice())
	// println("Price of veg pizza with cheese: ", cheesePizza.GetPrice())
	// println("Price of veg pizza with cheese and mushroom: ", mushroomPizza.GetPrice())

	// Adapter pattern
	client := &adapter.Client{}

	mac := &adapter.Mac{}
	client.ConnectToTypeC(mac)

	windowMachine := &adapter.Window{}
	windowsMachineAdapter := &adapter.WindowAdapter{
		Window: windowMachine,
	}

	client.ConnectToTypeC(windowsMachineAdapter)

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
