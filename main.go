package main

import (
	"fmt"
	// printingHelloworld "go-learnings/concepts/example-code/01-package-imports"
	controlstatementproblems "go-learnings/daily-progress/01-control-statements"
	arrays "go-learnings/daily-progress/02-arrays"
	DSA "go-learnings/DSA/Arrays"
)

func main() {

	// fmt.Println("Calling Other Methods Packages")
	// printingHelloworld.PrintHelloWorld()

	fmt.Println("========Control Statement Problems==================")
	controlstatementproblems.PrintPostiveNegativeNumbers(5)

	fmt.Println("========Array-Problems==================")
	arr := [5]int {1,2,3,4,5}
	arrays.ReverseArray(arr)

	convertingToJson()

	result := DSA.SecondLargest([]int{10,20,30,40,50})

	fmt.Println("Second Largest",result)

}
