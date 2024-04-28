package day11

import (
	"fmt"
	"sort"

	"aoc-2022/11/common"
	"aoc-2022/pkg/utils"
)

func SolveA() {
	lines := utils.ReadLines("11/input/real.txt")

	monkeysData := common.ParseInput(lines)
	monkeysData = common.InvestigateMonkeys(monkeysData, 3, 20)

	sort.Slice(monkeysData, func(i, j int) bool {
		return monkeysData[i].ItemCount > monkeysData[j].ItemCount
	})

	ans := monkeysData[0].ItemCount * monkeysData[1].ItemCount
	fmt.Println(ans)
}
