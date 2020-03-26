package main

import (
	"fmt"
)

func main() {
	var a int16
	a = 1

	var b int32 = 2

	c := 3 // type inferred

	fmt.Println(a) // 1
	fmt.Println(b) // 2
	fmt.Println(c) // 3

	fmt.Printf("%T\n", a) // int16
	fmt.Printf("%T\n", b) // int32
	fmt.Printf("%T\n", c) // int
}
