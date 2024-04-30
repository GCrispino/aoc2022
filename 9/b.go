package day9

import (
	"fmt"
)

func SolveB() {
	moves := parseMovesFromFile("9/input/a.txt")

	finalState := simulate(moves, 10)

	fmt.Println(finalState.visitedTracker.String())

	fmt.Println(len(finalState.visitedTracker))
}
