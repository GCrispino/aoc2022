package day1

import (
	"fmt"
	"strconv"

	"aoc-2022/pkg/utils"
)

func deersDataFromLines(lines []string) [][]string {
	deersData := make([][]string, 0)

	deersDataBegin := 0
	for i, line := range lines {
		if line == "" {
			deersData = append(deersData, lines[deersDataBegin:i])
			deersDataBegin = i + 1
		}
	}

	return deersData
}

func findMaxCaloriesDeer(deersData [][]string) (maxCaloriesDeer int) {
	for _, deerData := range deersData {
		sumCalories := 0
		for _, itemCal := range deerData {
			itemCalInt, err := strconv.Atoi(itemCal)
			if err != nil {
				panic(err)
			}

			sumCalories += itemCalInt
		}

		if sumCalories > maxCaloriesDeer {
			maxCaloriesDeer = sumCalories
		}
	}

	return
}

func SolveA() {
	lines := utils.ReadLines("1/input/real.txt")
	deersData := deersDataFromLines(lines)

	fmt.Println(findMaxCaloriesDeer(deersData))
}
