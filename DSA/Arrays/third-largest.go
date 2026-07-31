package Arrays

import "math"

func ThirdLargest(arr []int) int {

	firstMax := math.MinInt
	secondMax := math.MinInt
	thirdMax := math.MinInt

	for i := 0 ; i < len(arr) ; i++ {

		if arr[i] > firstMax {
			
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = arr[i]

		} else if arr[i] > secondMax && firstMax != arr[i]{

			thirdMax = arr[i]
			secondMax = arr[i]

		} else if arr[i] > thirdMax && arr[i] != secondMax && arr[i] != thirdMax {

			thirdMax = arr[i]

		}
	}

	return thirdMax
}