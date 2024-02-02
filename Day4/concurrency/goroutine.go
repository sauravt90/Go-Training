package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func Display(msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	fmt.Println("this is Printing")
	go fmt.Println("this is pringint in separate thread")

	wg.Add(1)
	go Display("jifisd ndng")
	wg.Add(1)
	go Display("eoke")

	fmt.Println("this is exiting")
	wg.Wait()
}
