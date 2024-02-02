package main

import "fmt"

type Employee struct {
	Empid int
	Name  string
}

// print
func (a Employee) Print() {
	fmt.Println("===========")
	fmt.Println("Employee Id -", a.Empid)
	fmt.Println("Employee Name -", a.Name)
	fmt.Println("===========")
}

//update

//call by value
func (a Employee) Update(name string) {
	a.Name = name
}

//call by ref
func (a *Employee) Update1(name string) {
	a.Name = name
}

func (a Employee) Initialize() {

}

func main() {
	e := Employee{23, "fee"}

	e.Print()
}
