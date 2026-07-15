package arrays
import "fmt"

func FindMinInArray(arr [5]int) {

	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	fmt.Println("Minimum Element in an array:",min)
	
}