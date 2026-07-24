package scope

import "fmt"

func CounterIncrease(n int) {
	counter:=0

	for i:=0;i<n;i++ {
		counter:=counter+1
		fmt.Println(counter)
	} // this is the bug
}

func CounterIncreaseProperly(n int) {
	counter:=0

	for i:=0;i<n;i++ {
		counter+=1
		fmt.Println(counter)
	}
	fmt.Println(counter)
}