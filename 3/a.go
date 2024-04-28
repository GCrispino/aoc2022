package day3

import (
	"fmt"

	"aoc-2022/pkg/utils"
	"aoc-2022/pkg/utils/types"

	"aoc-2022/3/common"
)

func getRucksacksData(lines []string) []types.Pair[string] {
	rucksacksData := make([]types.Pair[string], 0)

	for _, line := range lines {
		lenLine := len(line)
		half := lenLine / 2
		rucksacksData = append(rucksacksData, types.Pair[string]{
			First:  line[:half],
			Second: line[half:],
		})
	}

	return rucksacksData
}

func SolveA() {
	lines := utils.ReadLines("3/input/real.txt")
	rucksacksData := getRucksacksData(lines)

	total := 0

	for _, rData := range rucksacksData {
		commonVal := common.FindCommonItemStrings([]string{rData.First, rData.Second})

		total += common.CalcPriority(commonVal)
	}
	fmt.Println("Total:", total)

}
