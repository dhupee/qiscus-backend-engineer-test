package main

import (
	"bufio"
	"fmt"
	"os"
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

		// parse the line
		fmt.Println(line)

		// TODO: get the divisors thingy

		count++

		// close the file once you've finished
		defer file.Close()
	}
}
