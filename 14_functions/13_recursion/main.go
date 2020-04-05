package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	// don't do this at home!
	// for i := 1; i < 100; i++ {
	// 	fmt.Println("factorial", i, "=", factorial(i))
	// }
}
