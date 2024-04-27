package main

import (
  "fmt"
  "sort"
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

func getCaloriesFromDeersData(deersData [][]string) []int {
  deersCalories := make([]int, len(deersData))

  for i, deerData := range deersData {
    sumCalories := 0
    for _, itemCal := range deerData {
      itemCalInt, err := strconv.Atoi(itemCal)
      if err != nil {
        panic(err)
      }

      sumCalories += itemCalInt
    }

    deersCalories[i] = sumCalories
  }

  return deersCalories
}

func main(){
  lines := utils.ReadLines("1/input/real.txt")
  deersData := deersDataFromLines(lines)
  deersCalories := getCaloriesFromDeersData(deersData)
  sort.Ints(deersCalories)
  
  fmt.Println(utils.SumInts(deersCalories[len(deersCalories)-3:]))
}
