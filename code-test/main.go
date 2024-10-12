package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// save for later use
// type divisorsOfNumber struct {
// 	number       int
// 	divisorCount int
// 	evenDivisors bool
// }

func main() {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	numWithEvenDivisor := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// turn the line into an integer
		numInput, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		// count the divisors
		divisorCount := 0
		var evenDivisors bool
		for i := 1; i <= numInput; i++ {
			if numInput%i == 0 {
				divisorCount++
			}
			// find whether the number has an even number of divisors
			if divisorCount%2 == 0 {
				evenDivisors = true
			} else {
				evenDivisors = false
			}

		}

		// count how many numbers have an even number of divisors
		if evenDivisors {
			numWithEvenDivisor++
		}

		// print the result
		fmt.Println("Number:", numInput, "has", divisorCount, "divisors,", "even divisors is", evenDivisors)
		count++

		// close the file once you've finished
		defer file.Close()
	}

	fmt.Println("")
	fmt.Println("Number that has even number of divisors: ", numWithEvenDivisor)
}
