package arrays
import "fmt"

/* 
	Any changes you make to arr inside ReverseArray are only happening to that local copy.    The original array you passed in from main() will remain completely unchanged.

	To fix this, you should use a slice ([]int) instead of a fixed-size array ([5]int). Slices in Go act as references to the underlying array, meaning changes made inside the function will reflect on the original data.

	In coming problems we will use slice concept we will cover this topic next lessons with 
	detailed explaination.
*/

func ReverseArray(arr [5]int) {

	end := len(arr)-1

	for start := 0 ; start < end ; start++ {
		arr[start] , arr[end] = arr[end] , arr[start]
		end--
	}

	fmt.Println("Reverse the array:",arr)

}