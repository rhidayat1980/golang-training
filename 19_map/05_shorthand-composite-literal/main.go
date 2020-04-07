package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Time": "Good morning",
		"Jane": "Bonjour",
	}

	fmt.Println(myGreeting["Jane"])

	for k, v := range myGreeting {
		fmt.Println(k, v)
	}
}
