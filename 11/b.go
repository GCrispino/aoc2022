package day11

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"time"

	common "aoc-2022/11/common2"
	"aoc-2022/pkg/utils"
)

func SolveB() {
	nRounds := 15
	f, err := os.Create(fmt.Sprintf("%d-%d.prof", nRounds, time.Now().Unix()))
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	lines := utils.ReadLines("11/input/test.txt")

	fmt.Println("nRounds:", nRounds)

	monkeysData := common.ParseInput(lines)
	monkeysData = common.InvestigateMonkeys(monkeysData, 1, nRounds)

	sort.Slice(monkeysData, func(i, j int) bool {
		return monkeysData[i].ItemCount > monkeysData[j].ItemCount
	})

	fmt.Println(utils.Map(monkeysData, func(m *common.Monkey) int {
		return m.ItemCount
	}))

	ans := monkeysData[0].ItemCount * monkeysData[1].ItemCount
	fmt.Println(ans)
}
