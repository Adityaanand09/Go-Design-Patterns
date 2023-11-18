package FactoryPattern

type Daytime struct {
	Employee
}

func NewDayTimeEmp(n string, empType string) IEmployee {
	return &Daytime{Employee{
		name:    n,
		empType: empType,
	}}
}
