package Arrays

import "math"

/*
	Maximum product of a triplet (subsequence of size 3) in array
	Given an integer array, find a maximum product of a triplet in the array.
	Examples:

	Input:  arr[ ] = [10, 3, 5, 6, 20]
	Output: 1200
	Explanation: Multiplication of 10, 6 and 20

	Input:  arr[ ] =  [-10, -3, -5, -6, -20]
	Output: -90

	Input: arr[ ] =  [1, -4, 3, -6, 7, 0]
	Output: 168
*/

func maximumProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	firstMax := math.MinInt
	secondMax := math.MinInt
	thirdMax := math.MinInt

	firstMin, secondMin := math.MaxInt, math.MaxInt

	for i := 0; i < len(nums); i++ {

		if nums[i] > firstMax {

			thirdMax = secondMax
			secondMax = firstMax
			firstMax = nums[i]

		} else if nums[i] > secondMax {

			thirdMax = secondMax
			secondMax = nums[i]

		} else if nums[i] > thirdMax {
			thirdMax = nums[i]
		}

		if nums[i] < firstMin {

			secondMin = firstMin
			firstMin = nums[i]

		} else if nums[i] < secondMin {
			secondMin = nums[i]
		}
	}

	candidate1 := firstMax * secondMax * thirdMax
	candidate2 := firstMin * secondMin * firstMax

	if candidate1 > candidate2 {
		return candidate1
	} else {
		return candidate2
	}
}
