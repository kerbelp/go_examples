package binary_search

import "testing"

func TestBinary_search(t *testing.T) {
	slice := []int{1,2,3,4,5,6,7}
	num := 4
	
	pos := BinarySearch(num, slice, 0)
	if pos != 3 {
		t.Errorf("binary_search Failed to find number 4")
	}

	num2 := 9
	pos2 := BinarySearch(num2, slice, 0)
	if pos2 != -1 {
		t.Errorf("binary_search Failed to not find 9")	
	}
	
}
