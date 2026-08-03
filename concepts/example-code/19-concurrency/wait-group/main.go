package main

import (
	"fmt"
	"sync"
)

func sayHello(sync *sync.WaitGroup) {

	defer sync.Done()
	fmt.Println("Say Hello to world!")

}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)
	for i := 0 ; i < 3 ; i++ {
		go sayHello(&wg)
	}
	wg.Wait()
	fmt.Println("Main Function")
}