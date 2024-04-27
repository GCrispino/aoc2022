package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc-2022/pkg/utils"
	"aoc-2022/pkg/utils/types"
)

type assignment types.Pair[int]

func (a assignment) isContainedInto(into assignment) bool {
	return a.First >= into.First && a.Second <= into.Second
}

func (a assignment) overlapsWith(with assignment) bool {
	return a.First <= with.First && a.Second >= with.First
}


func NewAssignment(assignmentStr string) assignment {
	spl := strings.Split(assignmentStr, "-")
	begin, err := strconv.Atoi(spl[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(spl[1])
	if err != nil {
		panic(err)
	}

	return assignment{
		begin, end,
	}
}

type assignmentPair types.Pair[assignment]

func (a assignmentPair) hasContained() bool {
	first := a.First
	second := a.Second

	return first.isContainedInto(second) || second.isContainedInto(first)
}
 
func (a assignmentPair) hasOverlap() bool {
	if a.hasContained() {
		return true
	}

	first := a.First
	second := a.Second

	return first.overlapsWith(second) || second.overlapsWith(first)
}

func getAssignmentPairs(lines []string) []assignmentPair {
	assignmentPairs := make([]assignmentPair, len(lines))

	for i, line := range lines {
		spl := strings.Split(line, ",")
		assignmentPairs[i] = assignmentPair{
			First:  NewAssignment(spl[0]),
			Second: NewAssignment(spl[1]),
		}
	}

	return assignmentPairs
}

func main() {
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
