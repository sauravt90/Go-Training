package main

import (
	"fmt"
	"sync"
)

var count = 0

func main() {

	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := 0; idx < 10000; idx++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
