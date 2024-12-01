package main


func bubbleSort(list []int) {	
	for i := 0; i < len(list); i++ {
		for l := i; l < len(list); l++ {
			if list[i] > list[l] {
				list[i], list[l] = list[l], list[i]
			}
		}
	}
}

