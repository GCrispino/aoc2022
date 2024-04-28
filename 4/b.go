package day4

import (
	"fmt"

	"aoc-2022/pkg/utils"
)

func SolveB() {
	lines := utils.ReadLines("4/input/real.txt")
	assignmentPairs := getAssignmentPairs(lines)

	total := 0
	for _, pair := range assignmentPairs {
		if pair.hasOverlap() {
			total++
		}
	}

	// fmt.Println(assignmentPairs)

	fmt.Println(total)
}
