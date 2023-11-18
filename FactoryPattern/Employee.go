package FactoryPattern

type Employee struct {
	name    string
	empType string
}

func (e *Employee) GetEmpInfo() {
	println(e.name, e.empType)
}
