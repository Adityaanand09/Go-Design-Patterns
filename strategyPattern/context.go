package strategyPattern

type Context struct {
	v IVehicle
}

func (c *Context) SetVehicle(v IVehicle) {
	c.v = v
}

func (c *Context) Execute() {
	c.v.GetVehicleInfo()
	c.v.GetVehicleName()
}
