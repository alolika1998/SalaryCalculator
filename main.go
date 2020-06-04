package main

import "fmt"

type SalaryCalculator interface {
	CalculatorSalary() int
}
type Employee struct {
	Name string
	Salary int
	Pf int
}
type Stuff struct{
	Name string
	Salary int
}
func (e Employee) CalculatorSalary() int {
	return e.Salary + e.Pf
}
func (s Stuff) CalculatorSalary() int {
	return s.Salary
}
func TotalExpense(c []SalaryCalculator) {
Expense := 0
for _, v:= range c {
	Expense = Expense + v.CalculatorSalary()
}
fmt.Println("Total expense: ", Expense)
}
func main() {
	emp1 := Employee {
		Name: "Lisa",
		Salary: 6000,
		Pf: 1000,
	}
	emp2 := Employee {
		Name: "Sam",
		Salary: 8000,
		Pf:  1500,
	}
	st1 := Stuff {
		Name: "Maya",
		Salary: 3000,
	} 
	Team := []SalaryCalculator{emp1, emp2, st1}
	TotalExpense(Team)


}
