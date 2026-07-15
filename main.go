package main

import (
	"fmt"
	// printingHelloworld "go-learnings/concepts/example-code/01-package-imports"
	controlstatementproblems "go-learnings/daily-progress/control-statements"
	arrays "go-learnings/daily-progress/arrays"
)

func main() {

	// fmt.Println("Calling Other Methods Packages")
	// printingHelloworld.PrintHelloWorld()

	fmt.Println("========Control Statement Problems==================")
	controlstatementproblems.PrintPostiveNegativeNumbers(5)

	fmt.Println("========Array-Problems==================")
	arr := [5]int {1,2,3,4,5}
	arrays.ReverseArray(arr)

}
