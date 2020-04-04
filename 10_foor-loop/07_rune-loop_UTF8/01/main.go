package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)

	z := 10
	fmt.Println(z)
	b := string(z)
	fmt.Printf("%T \n", b)
}
