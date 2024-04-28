package day2

import (
	"fmt"

	"aoc-2022/pkg/utils"
)

type shape2 string

func (s shape2) getScore() int {
	switch s {
	case "A": // rock
		return 1
	case "B": // paper
		return 2
	case "C": // scissors
		return 3
	}

	return 0
}

type playData2 struct {
	opponentShape, yourShape shape2
}

func (p playData2) getOutcome() (outcome int) {
	switch p.yourShape {
	case "X": // lose
		outcome = 0
		switch p.opponentShape {
		case "A": // rock
			outcome += 3 // scissors
		case "B": // paper
			outcome += 1 // rock
		case "C": // scissors
			outcome += 2 // paper
		}
	case "Z": // win
		outcome = 6
		switch p.opponentShape {
		case "A": // rock
			outcome += 2 // paper
		case "B": // paper
			outcome += 3 // scissors
		case "C": // scissors
			outcome += 1 // rock
		}
	case "Y": // draw
		outcome = 3
		outcome += p.opponentShape.getScore()
	}

	return
}

func (p playData2) getScore() int {
	return p.yourShape.getScore() + p.getOutcome()
}

type gameData2 []playData2

func gameDataFromLines2(lines []string) gameData2 {
	playsData := make([]playData2, 0)

	for _, line := range lines {
		pData := playData2{
			yourShape:     shape2(line[2]),
			opponentShape: shape2(line[0]),
		}
		playsData = append(playsData, pData)
	}

	return playsData
}

func (g gameData2) getTotalScore() (totalScore int) {

	for _, pData := range g {
		totalScore += pData.getOutcome()
	}

	return
}

func SolveB() {
	lines := utils.ReadLines("2/input/real.txt")
	gameData := gameDataFromLines2(lines)

	fmt.Println(gameData.getTotalScore())
}
