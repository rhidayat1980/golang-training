package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("Have you no friends?")
	}
}
