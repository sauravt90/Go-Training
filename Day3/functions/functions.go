package main

import "fmt"

//Higer order function
func IterateSlice(arr []string, fn func(val string)) {
	for _, val := range arr {
		fn(val)
	}
}

func main() {

	//annonymous function calling immediatily
	func(msg string) {
		fmt.Println(msg)
	}("gjgjg")

	//annonymous function stornig in some veriable
	fn := func(msg string) {
		fmt.Println(msg)
	}
	//passing function to Higer order function
	IterateSlice([]string{"dfd", "Dfd", "dfs"}, fn)
	var a = 10
	IterateSlice([]string{"ff", "df", "Dffd"}, func(val string) {
		a++
		fmt.Println(val)
	})

	fmt.Println(a)

}
