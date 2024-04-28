package day6

import (
	"fmt"

	"aoc-2022/6/common"
	"aoc-2022/pkg/utils"
)

func SolveB() {
	lines := utils.ReadLines("6/input/real.txt")
	line := lines[0]

	i := common.FindMarker(line, 13)
	fmt.Println(i)
}
