package main

import "fmt"

func main() {
	var str = "This is new string"

	fmt.Println(str)

	var str1 = `this is new \n """"""
	tihs 
	iht
	tiht
	itwhh
	`
	fmt.Println(str1)

	for idx, val := range str1 {
		fmt.Printf("idx - %d value - %c\n", idx, val)
	}
	//string is immuatable

	// val1 := str1[0]

	// str1[0] = val1

}
