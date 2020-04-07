package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Ohayo!",
		"Selamat pagi!",
		"GUtten morgen!",
	}


	fmt.Println("greeting[1:2] = ", greeting[1:2])
	fmt.Println("greeting[:2] =", greeting[:2])
	fmt.Println("greeting[5:] =", greeting[5:])
	fmt.Println("greeting[:]: =", greeting[:])
}