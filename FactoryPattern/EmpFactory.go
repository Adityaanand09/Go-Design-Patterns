package FactoryPattern

func GetEmp(n string, t string) IEmployee {

	if t == "daytime" {
		return NewDayTimeEmp(n, t)
	}
	if t == "nighttime" {
		return NewNightTimeEmp(n, t)
	}
	println("Wrong Time")
	return nil
}
