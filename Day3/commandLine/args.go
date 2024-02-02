package main

import (
	"fmt"
	"os"
)

func main() {

	fs := os.DirFS(".")
	fmt.Printf("fs: %v\n", fs)
	// args := os.Args[1:]
	// fmt.Println(args)

	// count := flag.Int("count", 0, "represents total count")
	// name := flag.String("name", "", "name of argurmnets")

	// flag.Parse()

	// arr := flag.Args()

	// fmt.Println(*count)
	// fmt.Println(*name)

	// fmt.Println(arr)
}
