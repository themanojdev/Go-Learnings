package arrays

func RemoveDuplicates(nums []int) int {

	j:=0
	for i:=1;i<len(nums);i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return len(nums[:j+1])

}