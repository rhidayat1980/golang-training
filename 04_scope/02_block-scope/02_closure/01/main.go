package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	{
		fmt.Println(x)
		y := "The credit belongs to the one who is in the ring"
		fmt.Println(y)
	}

	// fmt.Println(y) is outside the scope, cannot run
}
