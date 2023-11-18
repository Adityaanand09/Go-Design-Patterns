package main

import (
	"DesignPatterns/strategyPattern"
	"DesignPatterns/strategyPattern/ConcreteStrategy"
)

func main() {

	bike := ConcreteStrategy.NewBike(250, "red", 100000, "Ninja")
	car := ConcreteStrategy.NewCar(2500, "Black", 1000000.0, 4, "BMW")

	context := &strategyPattern.Context{}

	context.SetVehicle(car)
	context.Execute()
	println()
	context.SetVehicle(bike)
	context.Execute()
}
