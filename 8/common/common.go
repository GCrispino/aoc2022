package common

import "aoc-2022/pkg/utils"

func LinesToMatrix(lines []string) [][]int {
	matrix := make([][]int, len(lines))

	for i, line := range lines {
		matrix[i] = utils.Map([]rune(line), func(x rune) int {
			return int(x) - '0'
		})
	}

	return matrix
}
