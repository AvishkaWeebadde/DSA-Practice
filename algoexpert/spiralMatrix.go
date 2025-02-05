package main

import (
	"fmt"
)

func SpiralTraverse(array [][]int) []int {
	x := len(array) - 1
    y := len(array[0]) - 1

	result := []int{}
	for i := 0; i <= y; i++ {
        result = append(result, array[0][i])
    }

	for i := 1; i <= y; i++ {
		result = append(result, array[i][y])
	}

	for i := y - 1; i >= 0; i-- {
		result = append(result, array[y][i])
	}

	for i := y - 1; i >= 1; i-- {
		result = append(result, array[i][0])
	}

	if x - 1 > 1 && y - 1 > 1 {
		newArray := [][]int{}
		for i := 1; i < x; i++ {
			newArray = append(newArray, array[i][1:y])
		}
		result = append(result, SpiralTraverse(newArray)...)
	}
	

	return result
}

func main() {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	fmt.Println(SpiralTraverse(array))
}