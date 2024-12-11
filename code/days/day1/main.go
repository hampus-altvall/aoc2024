package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers0 := []int{}
	numbers1 := []int{}
	// Create a scanner to read the file line-by-line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number_pair := strings.Split(line, "   ")

		num0, err := strconv.Atoi(number_pair[0])
		if err != nil {
			log.Fatal(err) // Handle error if the string is not a valid integer
		}
		num1, err := strconv.Atoi(number_pair[1])
		if err != nil {
			log.Fatal(err) // Handle error if the string is not a valid integer
		}

		numbers0 = append(numbers0, num0)
		numbers1 = append(numbers1, num1)

		// Print the converted integer
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Handle error if something goes wrong during scanning
	}

	sort.Ints(numbers0)
	sort.Ints(numbers1)

	length := len(numbers0)

	if length != len(numbers1) {
		log.Fatal("")
	}

	distance := 0
	for i := 0; i < length; i++ {
		a := numbers0[i]
		b := numbers1[i]

		if a > b {
			distance += a - b
		} else {
			distance += b - a
		}
	}

	fmt.Println(distance)
}
