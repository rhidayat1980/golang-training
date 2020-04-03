package main

import (
	"fmt"

	"github.com/rhidayat1980/golang-training/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	fmt.Println(vis.YourName)

	// print using an exported function inside vis package
	vis.PrintVar()
}

/*
to remove ugly import use go mod instead.
*/
