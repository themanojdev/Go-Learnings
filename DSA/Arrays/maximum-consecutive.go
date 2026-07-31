package Arrays

/*
	Given a binary array arr[] consisting of only 0s and 1s, find the length of the longest contiguous sequence of either 1s or 0s in the array.
*/
func MaxConsecBits(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	count,maxCount := 1,1

	for i := 0 ; i < len(nums)-1 ; i++ {

		if nums[i] == nums[i+1] {

			count++

		} else {
			count = 1
		}

		if count > maxCount {
			maxCount = count
		}

	}

	return maxCount

}