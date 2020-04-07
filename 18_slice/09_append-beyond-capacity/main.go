package main

import "fmt"

func main(){

	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"

	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zao 'an")
	greeting = append(greeting, "Ohayaou gozaimasu")
	greeting = append(greeting, "gidday")


	for _, v := range greeting {
		fmt.Println(v)
	}

	fmt.Println(greeting)
}
