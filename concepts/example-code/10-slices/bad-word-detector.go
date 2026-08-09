package slices

//O(n*m)
func indexOfFirstBadWordNavieApproach(msg,badWords []string) int {

	for i := 0 ; i < len(msg) ; i++ {
		for j := 0 ; i < len(badWords) ; j++ {
			if badWords[j] == msg[i] {
				return i
			}
		}
	}

	return -1
}

//O(n+m)
func indexOfFirstBadWordOptimizedApproach(msg,badWords []string) int {

	words := make(map[string]bool)

	for _,value := range badWords {
		words[value] = true
	}

	for i,value := range msg {
		if words[value] {
			return i
		}
	}

	return -1
}