package binary_search_with_position

func BinarySearchWithPosition(num int, slice []int, pos int) []int {
	mid := len(slice) / 2
	if mid == 0 && slice[mid] != num {
		return []int{-1,pos}
	}
	
	if slice[mid] == num {
		return []int{pos + mid, pos + mid}
	} 
	
	if slice[mid] > num {
		return BinarySearchWithPosition(num, slice[0:mid], pos)
	}
	
	if slice[mid] < num {
		return BinarySearchWithPosition(num, slice[mid: len(slice)], pos + mid)
	}
	
	return []int{-1,-1}	
}
