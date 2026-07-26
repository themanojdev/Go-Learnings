package arrays

import "math"

/*
thridLargestElement returns the third largest distinct element
from the given slice.

Approach:
  - Traverse the slice only once while maintaining the three
    largest distinct values found so far.
  - firstMax stores the largest element.
  - secMax stores the second largest distinct element.
  - thridMax stores the third largest distinct element.

Why math.MinInt?
  - All three variables are initialized with math.MinInt so the
    function works correctly for both positive and negative integers.

Return Value:
  - If three distinct elements exist, the function returns the
    third largest distinct element.
  - If fewer than three distinct elements are present, the function
    returns the largest element, matching the expected behaviour
    of the problem.
*/
func thridLargestElement(nums []int) int {

	firstMax := math.MinInt
	secMax := math.MinInt
	thridMax := math.MinInt

	for i := 0 ; i < len(nums) ; i++ {

		if nums[i] > firstMax {

			secMax = firstMax
			firstMax = nums[i]

		} else if nums[i] > secMax && nums[i] != firstMax {

			thridMax = secMax
			secMax = nums[i]
			
		} else if nums[i] > thridMax && nums[i] != firstMax && nums[i] != secMax {
			thridMax = nums[i]
		}
	}

	if thridMax == math.MinInt {

		thridMax = firstMax
		return thridMax

	} else {
		return thridMax
	}
}