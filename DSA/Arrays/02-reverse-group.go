package Arrays

/* 
	Given an array arr[] and an integer k, find the array after reversing every subarray of consecutive k elements in place. If the last subarray has fewer than k elements, reverse it as it is. Modify the array in place, do not return anything.
*/

func ReverseArrayGroupWise(arr []int, k int) {

	var start,end,n int = 0,k-1,len(arr)
	var groups = n/k

	for i:=1;i<=groups;i++ {

		reverse(start,end,arr)
		start += k
		end += k

	}

	if start < n-1 {
		reverse(start,n-1,arr)
	}

}

func reverse(start,end int,arr []int) {

	for start < end {

		arr[start],arr[end] = arr[end],arr[start]
		start++
		end--

	}

}
