package main

import (
	"fmt"
)

// interface can hold any type of argument as it has no methods in it (empty inteface) - think of comcept of interface to undestand
func Print(a ...interface{}) {
	for _, val := range a {
		fmt.Println(val)
	}
}

func main() {
	a := 2
	b := "dfae"
	k := 34.33
	c := false

	Print(a, b, c, k)

}
