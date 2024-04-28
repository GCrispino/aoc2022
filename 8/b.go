package day8

import (
	"fmt"

	"aoc-2022/8/common"
	"aoc-2022/pkg/utils"
)

func getScenicScores(m [][]int) [][]int {
	scenicScores := make([][]int, len(m))

	for i, row := range m {
		scenicScores[i] = make([]int, len(row))
		for j := range row {
			scenicScores[i][j] = getScenicScore(i, j, m)
		}
	}
	return scenicScores
}

func getScenicScore(i, j int, m [][]int) int {
	val := m[i][j]

	nRows := len(m)
	nCols := len(m[0])

	blockingUp := i
	blockingLeft := j
	blockingRight := nCols - j - 1
	blockingDown := nRows - i - 1

	// row
	for k, rowVal := range m[i] {
		if k == j {
			continue
		}

		if rowVal >= val {
			if k < j {
				blockingLeft = j - k
				// fmt.Println("  blockingLeft = ", blockingLeft)
			}
			if k > j && blockingRight == nCols-j-1 { // k > j
				blockingRight = k - j
				// fmt.Println("  blockingRight = ", blockingRight)
			}
		}
	}

	// column
	for k := 0; k < nCols; k++ {
		colVal := m[k][j]

		if k == i {
			continue
		}

		if colVal >= val {
			if k < i {
				blockingUp = i - k
				// fmt.Println("  blockingUp = ", blockingUp)
			}
			if k > i && blockingDown == nRows-i-1 { // k > i
				blockingDown = k - i
				// fmt.Println("  blockingDown = ", blockingDown)
			}
		}
	}

	scenicScore := blockingLeft * blockingRight * blockingDown * blockingUp

	// fmt.Printf(
	// 	"scenic score of tree of height %d at (%d, %d): %d (%d * %d * %d * %d)\n",
	// 	m[i][j], i, j, scenicScore,
	// 	blockingUp, blockingLeft, blockingRight, blockingDown,
	// )
	return scenicScore

}

func SolveB() {
	lines := utils.ReadLines("8/input/real.txt")

	treeMatrix := common.LinesToMatrix(lines)

	// for i := range treeMatrix {
	// 	fmt.Println(treeMatrix[i])
	// }

	scores := getScenicScores(treeMatrix)
	highest := 0
	for i := range scores {
		// fmt.Println(scores[i])
		for _, s := range scores[i] {
			if s > highest {
				highest = s
			}
		}
	}
	fmt.Println(highest)
}
