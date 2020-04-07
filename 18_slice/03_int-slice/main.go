package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)

	fmt.Println("----------------------")
	fmt.Println(mySlice)

	fmt.Println("length of mySlice =",len(mySlice))
	fmt.Println("capacity of mySlice =", cap(mySlice))


	fmt.Println("-----------------------")

	for i:=0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}

}