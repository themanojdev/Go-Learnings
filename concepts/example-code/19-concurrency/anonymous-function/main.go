package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0 ; i < 3 ; i++ {
		go func () {
			defer wg.Done()
			fmt.Print("I'm screct function")
		}()
	}

	wg.Wait()

	fmt.Println("Main function")
}
