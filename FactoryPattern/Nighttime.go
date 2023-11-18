package FactoryPattern

type Nighttime struct {
	Employee
}

func NewNightTimeEmp(n string, empType string) IEmployee {
	return &Nighttime{
		Employee: Employee{
			name:    n,
			empType: empType,
		},
	}

}
