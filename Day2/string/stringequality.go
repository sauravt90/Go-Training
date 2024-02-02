package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "Saurav Ga rier eirijwr"
	b := "Gaurav"
	arr := strings.Split(a, " ")
	if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is not equal to b")
	}

	fmt.Println(strings.Compare(a, b))

	arr2 := strings.Join(arr, `\`)

	fmt.Println(arr2)

}
