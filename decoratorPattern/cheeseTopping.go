package decorator_pattern

//This is the topping that has an association with the base class
//All the toppings have a IPizza member  since we will add the topping to the received pizza , the base pizza struct wont have any member, since it is the initial step to create the pizza

type CheeseTopping struct {
	Pizza IPizza
}

func (c CheeseTopping) GetPrice() int {
	price := c.Pizza.getPrice()
	return price + 17
}
