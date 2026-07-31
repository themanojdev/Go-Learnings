package Arrays

func AddingOneDigitToGivenNumber(nums []int) []int {

	result := 0
	place := 1
	
	for i := len(nums)-1 ; i >= 0 ; i-- {
		
		temp := nums[i]
		result = result + (temp*place)
		place *= 10
		
	}
  var slice []int

	for result > 0 {

		digits := result % 10
		slice = append([]int{digits},slice...)
		result /= 10

	}

	return slice
}