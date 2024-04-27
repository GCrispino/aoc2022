package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc-2022/pkg/utils"
)

func runProgram(lines []string) (totalSignalStrenght int){
	cycleNum := 1
	regXVal := 1

	adding := false
	iLine := 0
    signalStrenghts := make([]int, 0)
	for iLine < len(lines) {
		if cycleNum == 20 || (cycleNum-20)%40 == 0 {
            strength := cycleNum*regXVal
			signalStrenghts = append(signalStrenghts, strength)
            totalSignalStrenght += strength
			fmt.Println(" ", cycleNum, regXVal, signalStrenghts[len(signalStrenghts)-1])
		}
		line := lines[iLine]

		spl := strings.Split(line, " ")

		cmd := spl[0]

		if adding {
			addVal, err := strconv.Atoi(spl[1])
			if err != nil {
				panic(err)
			}

			regXVal += int(addVal)
			adding = false
			iLine++
		} else {
			if cmd == "addx" {
				adding = true
			} else {
				iLine++
			}
		}

		cycleNum++
	}

    return
}

func main() {
	lines := utils.ReadLines("10/input/real.txt")

    totalSignalStrenght := runProgram(lines)
    fmt.Println(totalSignalStrenght)
}
