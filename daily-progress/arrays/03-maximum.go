package arrays
import "fmt"

func FindMaxInArray(arr [5]int) {

	max := arr[0]
	for _,value := range arr {
		if max < value {
			max = value
		}
	}

	fmt.Println("Maximum Element in an array:",max)
	
}