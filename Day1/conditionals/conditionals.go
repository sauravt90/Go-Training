package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var weekDay int
	// fmt.Println("please enter the input")
	// fmt.Scanf("%d", &weekDay)
	// //switech case fallthrought keyword

	// switch weekDay {
	// case 0:
	// 	fmt.Println(("Sun"))
	// case 1:
	// 	fmt.Println(("Mon"))
	// case 2:
	// 	fmt.Println(("Tue"))
	// case 3:
	// 	fmt.Println(("Wed"))
	// case 4:
	// 	fmt.Println(("Thu"))
	// case 5:
	// 	fmt.Println(("Fri"))
	// case 6:
	// 	fmt.Println(("Sat"))
	// }

	// fmt.Println("please enter the input1")
	// var weekDay1 int
	// fmt.Scanf("%d", &weekDay1)
	// //switech case fallthrought keyword

	// switch weekDay1 {
	// case 0:
	// 	fmt.Println(("Sun"))
	// case 1:
	// 	fmt.Println(("Mon"))
	// case 2:
	// 	fmt.Println(("Tue"))
	// case 3:
	// 	fmt.Println(("Wed"))
	// case 4:
	// 	fmt.Println(("Thu"))
	// case 5:
	// 	fmt.Println(("Fri"))
	// case 6:
	// 	fmt.Println(("Sat"))
	// }

	var a, b, c = 100, 99, 99

	avg := (float64)(a+b+c) / float64(3)

	fmt.Println("avg", avg, reflect.TypeOf(avg))

	var slice = []int{2, 3, 4}
	slice = append(slice, 5, 6, 7)
	fmt.Println(slice)

	//inset element in slice at a perticular position
	//slice = append(slice[:3],)
}
