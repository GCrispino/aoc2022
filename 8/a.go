package day8

import (
	"fmt"

	"aoc-2022/pkg/utils"
	"aoc-2022/pkg/utils/types"

	"aoc-2022/8/common"
)

func getVisible(m [][]int) []types.Pair[int] {
	visibleCoords := make([]types.Pair[int], 0)

	for i, row := range m {
		for j := range row {
			if isVisible(i, j, m) {
				visibleCoords = append(visibleCoords, types.NewPair(i, j))
			}
		}
	}
	return visibleCoords
}

func isVisible(i, j int, m [][]int) bool {
	isVisibleRight := true
	isVisibleLeft := true
	isVisibleUp := true
	isVisibleDown := true

	val := m[i][j]

	// row
	for k, rowVal := range m[i] {
		if k == j {
			continue
		}

		if rowVal >= val {
			if k < j {
				isVisibleLeft = false
			} else { // k > j
				isVisibleRight = false
			}
		}
	}

	nCols := len(m[0])
	// column
	for k := 0; k < nCols; k++ {
		colVal := m[k][j]

		if k == i {
			continue
		}

		if colVal >= val {
			if k < i {
				isVisibleUp = false
			} else { // k > i
				isVisibleDown = false
			}
		}
	}

	return isVisibleRight || isVisibleLeft || isVisibleUp || isVisibleDown

}

func SolveA() {
	lines := utils.ReadLines("8/input/real.txt")

	treeMatrix := common.LinesToMatrix(lines)

	// for i := range treeMatrix {
	// 	fmt.Println(treeMatrix[i])
	// }

	visible := getVisible(treeMatrix)
	fmt.Println(len(visible))
}
