package main

import (
	"fmt"

	"aoc-2022/pkg/utils"
	"aoc-2022/5/common"
)


func main() {
	lines := utils.ReadLines("5/input/real.txt")
	problemData := common.GetProblemData(lines)

	problemData.ApplyPlan(common.S9001)
	fmt.Println(problemData.GetTop())
}
