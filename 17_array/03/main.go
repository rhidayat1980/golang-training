package main

import "fmt"

func main() {
	var x [256]int
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = i
	}

	for k, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if k > 50 {
			break
		}
	}

	fmt.Println(x)
}
