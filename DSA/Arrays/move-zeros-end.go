package Arrays

func MoveZerosToEnd(nums []int) {

	end := 0
	for start := 0 ; start < len(nums) ; start++ {

		if nums[start] != 0 {

			nums[end],nums[start] = nums[start],nums[end]
			end++

		}
	}

}