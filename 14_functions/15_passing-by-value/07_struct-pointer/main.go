package main

import "fmt"

type customer struct {
	name string
	age  int
}

func changeMe(z *customer) {
	fmt.Println("start")
	fmt.Println(z)
	fmt.Println(&z.name)
	z.name = "Rocky"
	fmt.Println(z)
	fmt.Println(&z.name)
	fmt.Println("End function")
}
func main() {
	customer1 := customer{"Todd", 40}
	fmt.Println(&customer1)

	changeMe(&customer1)

	fmt.Println(customer1)

	fmt.Println(&customer1.name)
	fmt.Println(&customer1.age)
}
