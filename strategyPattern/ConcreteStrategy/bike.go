package ConcreteStrategy

type Bike struct {
	engineCC int
	color    string
	price    int
	name     string
}

func NewBike(e int, c string, p int, n string) Bike {
	return Bike{engineCC: e, color: c, price: p, name: n}
}

func (b Bike) GetVehicleInfo() {
	println("Bike Info ->", b.price, b.color, b.engineCC)
}

func (c Bike) GetVehicleName() {
	println(c.name)
}
