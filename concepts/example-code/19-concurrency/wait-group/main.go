package main

import (
	"fmt"
	"sync"
)

type rect struct {
	width  int
	height int
}

var r = rect {
	width  : 10,
	height : 20,
}

func (r rect) Area() int {
	return r.height*r.width
}

type Person struct {
	name    string
	age     int
	reciver struct{
		address     string
		phoneNumber int
	}
}

var person = Person {
	name : "Mohith",
	age : 25,
	reciver: struct{address string; phoneNumber int}{
		address: "Bengaluru",
		phoneNumber: 9019107853,
	},
}

var Person2 = struct {
	name   string
	age    int
	marks  int
} {
	name : "Mohith",
	age : 24,
	marks : 100,
}

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

	fmt.Println(person.reciver.address)
}