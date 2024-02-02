package main

import "fmt"

func Variadic(str ...string) {
	for _, val := range str {
		fmt.Println(val)
	}
}

func main() {

	Variadic("sg")

	arr := []string{"G", "g"}
	Variadic(arr...)
}
