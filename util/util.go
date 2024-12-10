package util

import (
	"strconv"
	"strings"
)

// ReadInts will convert the line into a slice of ints
// helpful for lines like 1,2,3,4,5,6,7
func ReadInts(line string, sep string) ([]int, error) {
	var ints []int
	for _, s := range strings.Split(line, sep) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

// IsInBounds checks if the point (x, y) is within the bounds of the matrix
func IsInBounds[T int | string](twoDSlice [][]T, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(twoDSlice) && y < len(twoDSlice[0])
}

func Build2DMap(input []string) [][]string {
	m := make([][]string, len(input))
	for i, line := range input {
		m[i] = make([]string, len(line))
		for j, letter := range strings.Split(line, "") {
			m[i][j] = letter
		}
	}
	return m
}

func Build2DMapInt(input []string) [][]int {
	m := make([][]int, len(input))
	for i, line := range input {
		m[i] = make([]int, len(line))
		for j, numStr := range strings.Split(line, "") {
			num, _ := strconv.Atoi(numStr)
			m[i][j] = num
		}
	}
	return m
}
