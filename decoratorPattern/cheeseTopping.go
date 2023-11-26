package decorator_pattern

type CheeseTopping struct {
	Pizza IPizza
}

func (c CheeseTopping) GetPrice() int {
	price := c.Pizza.getPrice()
	return price + 17
}
