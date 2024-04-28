package day5

import (
	"fmt"

	"aoc-2022/5/common"
	"aoc-2022/pkg/utils"
)

func SolveB() {
	lines := utils.ReadLines("5/input/real.txt")
	problemData := common.GetProblemData(lines)

	problemData.ApplyPlan(common.S9001)
	fmt.Println(problemData.GetTop())
}
