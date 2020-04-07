package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// get the book adventures of sherlock holmes

	resp, err := http.Get("http://gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts

	buckets := make([]int, 12)

	// loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
