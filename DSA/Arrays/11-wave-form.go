package Arrays

func SortArrayInWave(nums []int) {

	for i := 1 ; i < len(nums) ; i+=2 {
		nums[i] , nums[i-1] = nums[i-1] , nums[i]
	}

}