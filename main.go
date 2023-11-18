package main

import "DesignPatterns/ObserverPattern"

func main() {

	customer1 := ObserverPattern.NewCustomer("xyz@f.com", "adi")
	customer2 := ObserverPattern.NewCustomer("123@5.com", "bdx")

	shirtItem := ObserverPattern.NewItem("Nike t shirt")
	shirtItem.Register(customer1)
	shirtItem.Register(customer2)
	shirtItem.UpdateAvailability()
}
