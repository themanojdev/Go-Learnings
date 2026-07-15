package arrays
import "fmt"

func SumOfArray(arr [5]int) {

	sum := 0
	for _, value := range arr {
		sum += value
	}

	fmt.Println("Sum of elements in an array:",sum)
	
}