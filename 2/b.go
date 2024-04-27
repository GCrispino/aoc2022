package main

import (
	"fmt"

	"aoc-2022/pkg/utils"
)

type shape string

func (s shape) getScore() int {
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

type playData struct {
	opponentShape, yourShape shape
}

func (p playData) getOutcome() (outcome int) {
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

func (p playData) getScore() int {
	return p.yourShape.getScore() + p.getOutcome()
}

type gameData []playData

func (g gameData) getTotalScore() (totalScore int) {

	for _, pData := range g {
		totalScore += pData.getOutcome()
	}

	return
}

func gameDataFromLines(lines []string) gameData {
	playsData := make([]playData, 0)

	for _, line := range lines {
		pData := playData{
			yourShape:     shape(line[2]),
			opponentShape: shape(line[0]),
		}
		playsData = append(playsData, pData)
	}

	return playsData
}

func main() {
	lines := utils.ReadLines("2/input/real.txt")
	gameData := gameDataFromLines(lines)

	fmt.Println(gameData.getTotalScore())
}
