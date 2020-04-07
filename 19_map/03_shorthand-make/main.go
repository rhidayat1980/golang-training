package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)
	myGreeting["Time"] = "Good morning"
	myGreeting["Jenny"] = "Hello"

	fmt.Println(myGreeting)
}
