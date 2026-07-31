package Arrays

import "math"

func SecondLargest(arr []int) int {

	firstMax := math.MinInt
	secondMax := math.MinInt
	
	for i:=0;i<len(arr);i++ {

		if arr[i] > firstMax {

			
			secondMax = firstMax
			firstMax = arr[i]

		} else if arr[i] > secondMax && arr[i] != firstMax {

			secondMax = arr[i]

		}

	}
	

	return secondMax
}