package main

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fib(i))
	}

}
