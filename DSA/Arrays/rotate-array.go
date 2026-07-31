package Arrays


// given array -> [1, 2, 3, 4, 5, 6] expected array -> [3, 4, 5, 6, 1, 2]

// 1st -> [2,1,3,4,5,6] index {0,d-1,arr}
// 2nd -> [2,1,6,5,4,3] index {d,n-1,arr}
// 3rd -> [3,4,5,6,1,2] correct

/*
	Rotate an Array by d - Counterclockwise or Left
	Given an array of integers arr[] of size n, the task is to rotate the array elements to the left by d positions.
*/

func RotateArray(arr []int,d int) {

	var start, n int = 0, len(arr)-1

 reverse(start,d-1,arr)
 reverse(d,n,arr)
 reverse(start,n,arr)

}