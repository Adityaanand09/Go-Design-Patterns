package decorator_pattern

type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.Pizza.getPrice()
	return pizzaPrice + 10
}
