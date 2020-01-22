package main

import "fmt"

func main() {
	fib := fibgen()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
func fibgen() func() int {
	a := 0
	b := 1
	return func() int {
		b, a = (a + b), b
		return b
	}
}
