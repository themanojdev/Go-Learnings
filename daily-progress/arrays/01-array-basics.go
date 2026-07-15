package arrays
import "fmt"

func PrintArrays(arr [5]int) {

	for _ , value := range arr {

		fmt.Println("Printing the elements in an array:",value)
	}
	
}