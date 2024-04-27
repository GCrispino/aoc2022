package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc-2022/pkg/utils"
)

type crtScreen []rune  

func (c crtScreen) String() (s string) {
	for i, r := range c {
		if i % 40 == 0 {
			s += "\n"
		}
		s += string(r)
	}

	return s
}

func (c *crtScreen) draw(spritePos, cycleNum int){
	cycleNum = cycleNum % 40
	toDraw := '.'
	fmt.Println("  draw", cycleNum, spritePos)
	if cycleNum >= spritePos && cycleNum <= spritePos + 2 {
		toDraw = '#'
	}

	*c = append(*c, toDraw)
}

func runProgram(lines []string) crtScreen {
	screen := make(crtScreen, 0)

	cycleNum := 1
	regXVal := 1

	adding := false
	iLine := 0

	for iLine < len(lines) {
		fmt.Println("  line no ", iLine + 1)
		screen.draw(regXVal, cycleNum)
		fmt.Println(screen)

		// if cycleNum == 50 {
		// 	break
		// }

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

    return screen
}

func main() {
	lines := utils.ReadLines("10/input/test.txt")

    screen := runProgram(lines)
    fmt.Println(screen)
}
