package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo file List")
	dir := flag.String("name", "./", "root directory to list files")
	fmt.Println("dir", *dir)
	flag.Parse()
	fmt.Println("dir after", *dir)
	list, err := os.ReadDir(*dir)

	if err != nil {
		fmt.Println("Recieved Error while reading directiry ")
		os.Exit(1)
	}
	for _, val := range list {
		fmt.Println(val.Name())
	}
}

//intialize module fo mod inint name of module

//
