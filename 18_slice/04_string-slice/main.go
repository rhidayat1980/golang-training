package main

import "fmt"

func main(){
	greeting:= []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}
	fmt.Println("------------------")
	for j:= 0; j < len(greeting); j++{
		fmt.Println(greeting[j])
	}
	fmt.Println("------------------")
	for _, values := range greeting {
		fmt.Println(values)
	}
}
