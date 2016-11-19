package binary_search_with_position

import "testing"

func TestBinary_search(t *testing.T) {
	slice := []int{1,2,3,4,5,6,7}
	num := 4
	
	pos := BinarySearchWithPosition(num, slice, 0)
	if pos[0] != 3 {
		t.Errorf("binary_search Failed to find number 4")
	}

	num2 := 9
	pos2 := BinarySearchWithPosition(num2, slice, 0)
	if pos2[0] != -1 && pos[1] != 6 {
		t.Errorf("pos2 is %v. expected: %v", pos2)	
	}
	
}
