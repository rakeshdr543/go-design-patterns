package main

import (
	factory "github.com/rakeshdr543/go-design-patterns/factory"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	println("Gun: ", g.GetName())
	println("Power: ", g.GetPower())
	println()
}
