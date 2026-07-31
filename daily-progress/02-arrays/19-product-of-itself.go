package arrays

//given array -> [1,2,3,4], prefix = 6
//result array ->  [1,1,2,6]

//prefix[] , suffix[] --> [prefix[] * suffix[]] = product of number itself
/*
	suffix = 24
	result[1,1,2,6] 
	nums[i] = [1,2,3,4]
	
	1st -> [1,1,2,1]
	2nd -> [1,1,4,1]
	3rd -> [1,12,4,1]
	4th -> [24,12,4,1]
*/
func ProductOfItself(nums []int) []int {

	result := make([]int,len(nums))

	prefix :=1

	for i:=0;i<len(nums);i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	
	for i:= len(nums)-1;i>=0;i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result

}