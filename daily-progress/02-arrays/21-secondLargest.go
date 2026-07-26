package arrays

/*
secondLargest returns the second largest distinct element
from the given slice.

Assumptions:
  - This implementation only works correctly for non-negative integers.

Limitations:
  - secLargest is initialized to -1, so it will not work
    correctly if the slice contains negative numbers.
  - To support negative integers, initialize both
    largest and secLargest with math.MinInt.

Return Value:
  - Returns the second largest distinct element.
  - Returns -1 if no second largest element exists.
*/
func secondLargest(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	largest := nums[0]
	secLargest := -1

	for i := 1 ; i < len(nums) ; i++ {

		if nums[i] > largest {

			secLargest = largest
			largest = nums[i]

		} else if nums[i] > secLargest && nums[i] != largest {
			secLargest = nums[i]
		}

	}

	return secLargest
}