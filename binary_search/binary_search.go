package binary_search

func BinarySearch(num int, slice []int, pos int) int {
	mid := len(slice) / 2
	if mid == 0 && slice[mid] != num {
		return -1
	}
	
	if slice[mid] == num {
		return pos + mid
	} 
	
	if slice[mid] > num {
		return BinarySearch(num, slice[0:mid], pos)
	}
	
	if slice[mid] < num {
		return BinarySearch(num, slice[mid: len(slice)], pos + mid)
	}
	
	return -1	
}
