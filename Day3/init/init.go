package main

import "fmt"

func init() {
	fmt.Println("calling my init function")
}

func init() {
	fmt.Println("second init")
}

func main() {
	fmt.Print("this is from main")
}
