package Arrays

func ReverseArray(arr []int)  {

	var start,end int = 0,len(arr)-1

	for start < end {

		arr[start],arr[end] = arr[end],arr[start]
		start++
		end--

	}

}