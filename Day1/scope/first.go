package main

import (
	"fmt"
)

var global int

func main() {
	fmt.Println("In main .....")
	fmt.Println("Value of global veriable is ", global)
	global = 100
	mySetter()
	fmt.Println("Value of global veriable after mySetter() ", global)

}
