package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type divisorsOfNumber struct {
	number       int
	divisorCount int
	evenDivisors bool
}

func main() {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

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
			if divisorCount%2 == 0 {
				evenDivisors = true
			} else {
				evenDivisors = false
			}
		}

		count++

		fmt.Println("Number: ", numInput, "has", divisorCount, "divisors", "even divisors:", evenDivisors)

		// close the file once you've finished
		defer file.Close()
	}
}
