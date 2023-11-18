package ConcreteStrategy

type Car struct {
	engineCC int
	color    string
	price    float32
	seats    int
	name     string
}

func NewCar(e int, c string, p float32, s int, n string) Car {
	return Car{engineCC: e, color: c, price: p, seats: s, name: n}
}

func (c Car) GetVehicleInfo() {
	println("Car Info ->", c.price, c.color, c.engineCC)
}

func (c Car) GetVehicleName() {
	println(c.name)
}
