package main

import "fmt"

func BinarySearch(num int, slice []int, pos int) int {
	fmt.Printf("checking for %d in %v. current position: %d\n", num, slice, pos)	
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

func main() {
	slice := []int{2, 3, 5, 7, 11, 13}
	num := int(4)
	
	pos := BinarySearch(num, slice, 0)
	if pos == -1 {
		fmt.Printf("could not find number %d in %v\n", num, slice)		
	} else {
		fmt.Printf("number %d found in position: %d\n", num, pos)
	}
}
