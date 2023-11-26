package main

import (
	decorator_pattern "DesignPatterns/decoratorPattern"
	"fmt"
)

func main() {
	veggieLover := decorator_pattern.VeggieMania{}

	cheeseTopping := decorator_pattern.CheeseTopping{Pizza: veggieLover}
	fmt.Printf("Price = %d", cheeseTopping.GetPrice())
}
