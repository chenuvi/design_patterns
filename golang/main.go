package main

import (
	"fmt"
	factory "golang/create_mode/factory_method"
)

func main() {
	fmt.Print("------main run--------")
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
