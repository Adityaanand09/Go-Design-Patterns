package decorator_pattern

//This is the base class that is going to be decorated later on with toppings
//This class implements the getPrice method

type VeggieMania struct {
}

func (v VeggieMania) getPrice() int {
	return 15
}
