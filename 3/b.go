package day3

import (
	"fmt"

	"aoc-2022/pkg/utils"

	"aoc-2022/3/common"
)

func splitRucksackGroups(lines []string) [][]string {
	lenLines := len(lines)
	nGroups := lenLines / 3

	groups := make([][]string, nGroups)

	for i := 0; i < nGroups; i++ {
		k := i * 3
		groups[i] = lines[k : k+3]
	}

	return groups
}

func SolveB() {
	lines := utils.ReadLines("3/input/real.txt")

	groups := splitRucksackGroups(lines)

	total := 0
	for _, group := range groups {
		commonVal := common.FindCommonItemStrings(group)
		total += common.CalcPriority(commonVal)

	}

	fmt.Println("Total:", total)
}
