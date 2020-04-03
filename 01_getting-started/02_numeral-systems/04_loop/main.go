package main

import "fmt"

func main() {

	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("decimal: %d \t binary: %b \t hexa: %x \n", i, i, i)
	}

	for i := 0; i <= 100; i++ {
		fmt.Printf("decimal: %d \t binary: %b \t hexa: %x \n", i, i, i)
	}

	num := 0
	for num < 10 {
		fmt.Println(num)
		num = num + 1
	}
	fmt.Println("last num:", num)
}

// disini kita diminta untuk nge print nilai desimal, binary dan hexa dari sebuah range value antara
// 1000000 - 1000100
