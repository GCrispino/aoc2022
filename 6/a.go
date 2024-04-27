package main

import (
	"fmt"

	"aoc-2022/pkg/utils"
	"aoc-2022/6/common"
)

func main() {
	lines := utils.ReadLines("6/input/real.txt")
	line := lines[0]

	i := common.FindMarker(line, 3)
    fmt.Println(i)
}
