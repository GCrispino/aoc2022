package day12

import (
	"aoc-2022/pkg/utils"
	"fmt"
	"math"
)

func getMinimumCostLowestElevation(grid [][]string, values [][]int) int {
	cost := math.MaxInt32
	for i := range grid {
		for j := range grid[i] {
			elevation := grid[i][j]
			if elevation != "a" {
				continue
			}
			value := values[i][j]

			if value < cost {
				cost = value
			}
		}
	}

	return cost
}

func SolveB() {
	path := "12/input/real.txt"
	lines := utils.ReadLines(path)

	state := NewState(lines)

	state.findOptimalPlan()

	minCostLowestElevation := getMinimumCostLowestElevation(state.grid, state.values)

	fmt.Println(minCostLowestElevation)
}
