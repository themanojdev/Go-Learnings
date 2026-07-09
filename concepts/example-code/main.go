package main

import (
	"fmt"
	printingHelloworld "go-learnings/concepts/example-code/01-package-imports"
)

func main() {
	fmt.Println("Calling Other Methods Packages")
	printingHelloworld.PrintHelloWorld()
}
