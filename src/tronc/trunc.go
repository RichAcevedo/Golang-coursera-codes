package main

import (
	"fmt"
)

func main() {
	var x float32

	fmt.Printf("insert a floating number: ")

	fmt.Scan(&x)
	y := int(x)
	fmt.Println("the number should be ", y)

}
