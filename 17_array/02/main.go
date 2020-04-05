package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
	fmt.Println(x[57])

	for k, v := range x {
		fmt.Println("index ", k, "value =", v)
	}

	for i := 65; i <= 122; i++ {
		fmt.Printf("%d decimal = %c in char/run = %x in hexa = %b in binary\n", i, i, i, i)
	}
}
