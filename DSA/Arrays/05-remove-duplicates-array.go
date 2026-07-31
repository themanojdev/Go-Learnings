package Arrays
// It will only wrok for sorted array.	
func RemoveDuplicatesFromArray(nums []int) []int {

	j := 0
	for i := 1 ; i < len(nums) ; i++ {

		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return nums[:j+1]

}