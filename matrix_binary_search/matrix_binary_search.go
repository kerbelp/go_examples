package main

import (
	"fmt"
	
	"github.com/kerbelp/go_examples/binary_search_with_position"
)

/*
	MatrixBinarySearch will search for a number in a n*m matrix.
	Algorythm is:
		1. binary search on first column - find suitable line
		2. binary search on single line 
*/
func MatrixBinarySearch(matrix [][]int, num int) []int{
	fmt.Printf("Looking for %d in %v\n", num, matrix)
	
	m := len(matrix)
	n := len(matrix[0])
	
	fmt.Printf("matrix is: %dX%d\n",n,m)
		
	firstColumn := []int{}
	for _, v := range matrix {
		firstColumn = append(firstColumn, v[0])
	}
	
	fmt.Printf("first column: %v\n", firstColumn)	
	rowPos := binary_search_with_position.BinarySearchWithPosition(num, firstColumn, 0)
	
	if rowPos[0] != -1 {
		return []int{rowPos[0], 0}
	}	
	fmt.Printf("number is in row: %d\n", rowPos[1])
	
	colPos := binary_search_with_position.BinarySearchWithPosition(num, matrix[rowPos[1]], 0)
		
	if colPos[0] == -1 {
		return []int{-1,-1}
	} else {
		return []int{rowPos[1], colPos[1]}
	}
	
	return []int{-1,-1}
}

func main() {
	matrix := [][]int{
		{2, 3, 5},  
		{7, 11, 14}, 
		{20, 22, 24},
		{25, 27, 31},
		}
	num := int(14)

	pos := MatrixBinarySearch(matrix, num)
	if pos[0] == -1 {
		fmt.Printf("Could not find %d\n", num)
	} else {
		fmt.Printf("Found it in position %v\n", pos)
	}
}






