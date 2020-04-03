package main

import "fmt"

func main() {
	number := []int{20, 30, 40, 50}
	for _, value := range number {
		fmt.Printf("decimal value: %d \t hexadecimal value: %#X \n", value, value)
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("decimal value: %d \t binary value: %b \t hexa value: %x\n", i, i, i)
	}
}
