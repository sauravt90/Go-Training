package main

import "fmt"

type Address struct {
	city    string
	adLine1 string
}

type Employee struct {
	name    string
	id      int
	address Address
}

func main() {
	e := Employee{name: "saurav", id: 2323}
	var emp Employee = Employee{name: "Gaurav", id: 07}
	e.name = "DADA"

	address := Address{"banglore", "nagwara"}

	e.address = address
	fmt.Println(e, emp)
}
