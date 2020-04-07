package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick

	res, err := http.Get("http://gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create a slice to hold counts
	buckets := make([]int, 200)

	// loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])

	fmt.Println("****************************")

	for i := 28; i <= 126; i++ {
		fmt.Printf("%v - %c - %v \n", i, i, i)
	}
}

func hashBucket(word string) int {
	return int(word[0])
}
