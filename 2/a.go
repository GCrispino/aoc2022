package main

import (
	"fmt"

	"aoc-2022/pkg/utils"
)

// type shape string
type shape int

const (
	Rock = iota
	Paper
	Scissors
)

func (s shape) getScore() int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	return 0
}

func NewShape(id string) shape {
	switch id {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	return Rock
}

type playData struct {
	opponentShape, yourShape shape
}

var scoreMap map[playData]int = map[playData]int{
	{yourShape: Rock, opponentShape: Rock}:     3,
	{yourShape: Rock, opponentShape: Paper}:    0,
	{yourShape: Rock, opponentShape: Scissors}: 6,
	//
	{yourShape: Paper, opponentShape: Paper}:    3,
	{yourShape: Paper, opponentShape: Rock}:     6,
	{yourShape: Paper, opponentShape: Scissors}: 0,
	//
	{yourShape: Scissors, opponentShape: Scissors}: 3,
	{yourShape: Scissors, opponentShape: Paper}:    6,
	{yourShape: Scissors, opponentShape: Rock}:     0,
}

func (p playData) getOutcome() int {
	return scoreMap[p]
}

func (p playData) getScore() int {
	fmt.Println(p.yourShape, p.opponentShape, p.yourShape.getScore(), scoreMap[p])
	return p.yourShape.getScore() + p.getOutcome()
}

type gameData []playData

func (g gameData) getTotalScore() (totalScore int) {

	for _, pData := range g {
		totalScore += pData.getScore()
	}

	return
}

func gameDataFromLines(lines []string) gameData {
	playsData := make([]playData, 0)

	for _, line := range lines {
		pData := playData{
			yourShape:     NewShape(line[2:3]),
			opponentShape: NewShape(line[0:1]),
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
