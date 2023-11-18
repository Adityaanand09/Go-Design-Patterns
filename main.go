package main

import "DesignPatterns/FactoryPattern"

func main() {

	emp1 := FactoryPattern.GetEmp("abc", "daytime")
	emp1.GetEmpInfo()

	emp2 := FactoryPattern.GetEmp("def", "nighttime")
	emp2.GetEmpInfo()
}
