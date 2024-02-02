package main

import "fmt"

//concurrency
//reflection
//interface
func Divide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func main() {
	defer func() {
		resc := recover()
		if resc != nil {
			fmt.Println("error whle executing")
		}

	}()

	res := Divide(3, 3)
	fmt.Println(res)

}
