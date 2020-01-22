package main

import (
	"fmt"
)

//func main() {
	compute := add()
	fmt.Println(compute())
}
func add() func() int {

	a := 12
	b := 34
	return func() int {
		c := a + b
		fmt.Println("closure add fx")
		return c
	}
}
