package main

import "fmt"

func main() {

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("your name either Marcus os Medhi")
	case myFriendsName == "Julian":
		fmt.Println("wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("wassup Sushant")
	default:
		fmt.Println("nothing matched; this is default")
	}
}
