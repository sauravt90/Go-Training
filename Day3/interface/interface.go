package main

import "fmt"

type Test interface {
	Print()
}

type Person struct {
	id   int
	name string
}

func (p Person) Print() {
	fmt.Println(p.name, p.id)
}
func (e Employee) Print() {
	fmt.Println(e.id, e.position)
}

type Employee struct {
	id       int
	position string
}

func main() {

	var obj Test

	p := Person{3, "saruav"}

	obj = p
	obj.Print()

	e := Employee{3, "engineer"}
	obj = e
	obj.Print()

}
