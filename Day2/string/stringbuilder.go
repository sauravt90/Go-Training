package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.Write([]byte("Hello"))
	builder.WriteString("12345")
	builder.WriteRune('$')

	str := builder.String()

	fmt.Println(str)
}
