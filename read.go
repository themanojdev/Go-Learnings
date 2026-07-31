package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define struct named name
type name struct {
	fname string
	lname string
}

func readThroughFile() {
	var filename string
	
	fmt.Print("Enter the name of the file: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []name

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if len(parts) >= 2 {
			person := name{
				fname: parts[0],
				lname: parts[1],
			}

			names = append(names, person)
		}
	}

	// Check scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, person := range names {
		fmt.Println(person.fname, person.lname)
	}
}