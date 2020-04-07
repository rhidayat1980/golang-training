package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)

	myGreeting["Time"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

	fmt.Println(myGreeting["Time"])

	for k, v := range myGreeting {
		fmt.Println(k, v)
	}

	for k := range myGreeting {
		fmt.Println(k)
	}

	for _, v := range myGreeting {
		fmt.Println(v)
	}

	fmt.Println(myGreeting)
}
