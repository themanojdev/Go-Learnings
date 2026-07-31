package Arrays

import "math"
/*
Given an integer array nums, find a subarray that has the largest product, and return the product.

A subarray is a contiguous non-empty sequence of elements within an array.

You can assume the output will fit into a 32-bit integer.

Note that the product of an array with a single element is the value of that element.

Example 1:

Input: nums = [2,4,-3,5]

Output: 8
Explanation: [2,4] has the largest product 8.

Example 2:

Input: nums = [-3,0,-2]

Output: 0
*/

func MaximumProductSubarray(nums []int) int {

	prefix,suffix,n := 1,1,len(nums)

	maxProduct := math.MinInt

	for i := 0 ; i < n ; i++ {
		
		if prefix == 0 {
			prefix = 1
		}

		if suffix == 0 {
			suffix = 1
		}
		
		prefix *= nums[i]
		if prefix > maxProduct {
			maxProduct = prefix
		}

		suffix *= nums[n-i-1]
		if suffix > maxProduct {
			maxProduct = suffix
		}

	}

	return maxProduct
}

