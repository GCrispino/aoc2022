package day12

import (
	"fmt"
	"strings"

	"aoc-2022/pkg/utils"
)

type position struct {
	i, j int
}

type state struct {
	grid            [][]string
	initialStatePos position
	goalPos         position

	graph map[position]neighborStateData

	values [][]int
}

type neighborStateData map[position]action

type action string

const (
	ActionLeft  action = "<"
	ActionRight action = ">"
	ActionUp    action = "^"
	ActionDown  action = "v"
)

func (s state) gridString() string {
	strLines := make([]string, len(s.grid))
	for i, line := range s.grid {
		strLines[i] = strings.Join(line, "")
	}

	return strings.Join(strLines, "\n")
}

func (s state) findOptimalPlan() [][]action {
	gridWidth := len(s.grid)
	gridLength := len(s.grid[0])

	optimalPlanGrid := make([][]action, gridWidth)
	for i := 0; i < gridWidth; i++ {
		optimalPlanGrid[i] = make([]action, gridLength)
	}

	maxIterations := len(s.graph)
	for n := 0; n < maxIterations; n++ {

		var avgChange float64

		changed := false
		for pos, neighborData := range s.graph {
			if pos == s.goalPos {
				continue
			}
			var minVal *int
			var minAction action
			for neighborPos, action := range neighborData {
				neighborVal := s.values[neighborPos.i][neighborPos.j]
				cost := 1 + neighborVal
				if minVal == nil {
					minVal = new(int)
					*minVal = cost
					minAction = action
					// } else if cost < s.values[pos.i][pos.j] {
				} else if cost < *minVal {
					*minVal = cost
					minAction = action
				}
			}
			// if minVal == nil {
			// 	continue
			// }
			if *minVal != s.values[pos.i][pos.j] {
				if n > 1000 {
					// fmt.Printf("    pos (%d, %d) value changed from %d to %d\n", pos.i, pos.j, s.values[pos.i][pos.j], *minVal)
				}
				avgChange = avgChange + float64(*minVal-s.values[pos.i][pos.j])
			}
			if minAction != optimalPlanGrid[pos.i][pos.j] {
				changed = true
				if n > 1000 {
					// fmt.Printf("    pos (%d, %d) value changed from %d to %d\n", pos.i, pos.j, s.values[pos.i][pos.j], *minVal)
					// fmt.Printf("    pos (%d, %d) action changed from %s to %s\n", pos.i, pos.j, optimalPlanGrid[pos.i][pos.j], minAction)
				}
			}
			s.values[pos.i][pos.j] = *minVal
			optimalPlanGrid[pos.i][pos.j] = minAction
		}

		avgChange = avgChange / float64(len(s.graph))
		// if n%100 == 0 {
		fmt.Printf("  iteration = %d, avgChange: %.5f, value at initial state: %d\n", n, avgChange, s.values[s.initialStatePos.i][s.initialStatePos.j])
		// }

		if !changed {
			break
		}
	}

	return optimalPlanGrid
}

func findReachableFromGraph(initialStatePos position, graph map[position]neighborStateData, visited map[position]struct{}) map[position]struct{} {
	if visited == nil {
		visited = make(map[position]struct{})
	}

	visited[initialStatePos] = struct{}{}
	for neighbor := range graph[initialStatePos] {
		_, ok := visited[neighbor]
		if !ok {
			visited = findReachableFromGraph(neighbor, graph, visited)
		}
	}

	return visited
}

func (s *state) simulatePlan(plan [][]action, logSteps bool) (bool, int) {
	curPosition := s.initialStatePos
	iStep := 0
	for curPosition != s.goalPos {
		action := plan[curPosition.i][curPosition.j]
		newPosition := curPosition
		switch action {
		case ActionRight:
			newPosition.j = curPosition.j + 1
		case ActionLeft:
			newPosition.j = curPosition.j - 1
		case ActionDown:
			newPosition.i = curPosition.i + 1
		case ActionUp:
			newPosition.i = curPosition.i - 1
		default:
			panic("invalid action")
		}

		// verify if there's better action
		// neighborStateData := s.graph[curPosition]
		curPosValue := s.values[curPosition.i][curPosition.j]
		possibleNeighbors := getPossibleNeighbors(curPosition, s.grid)
		// for pos := range neighborStateData {
		for _, pos := range possibleNeighbors {
			neighborValue := s.values[pos.i][pos.j]
			if 1+neighborValue < curPosValue {
				fmt.Printf("there might be something wrong with the value %d of position %v at elevation %s because of neighbor at %v with value %d and elevation %s!\n", curPosValue, curPosition, s.grid[curPosition.i][curPosition.j], pos, neighborValue, s.grid[pos.i][pos.j])
			}
		}

		if logSteps {
			fmt.Printf(" step: %d, curPosition: %v, elevation: %s, action: %s, newPosition: %v, new elevation: %s\n", iStep, curPosition, s.grid[curPosition.i][curPosition.j], action, newPosition, s.grid[newPosition.i][newPosition.j])
		}
		curPosition = newPosition
		iStep += 1
	}

	if curPosition == s.goalPos {
		return true, iStep
	}
	return false, iStep
}

func NewState(lines []string) state {
	nLines := len(lines)
	nColumns := len(lines[0])
	values := make([][]int, nLines)

	grid := make([][]string, nLines)

	var initialStatePos, goalPos position
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
		values[i] = make([]int, nColumns)
		for j, c := range grid[i] {
			statePos := position{i, j}
			if string(c) == "S" {
				initialStatePos = statePos

				grid[i][j] = "a"
			}
			if string(c) == "E" {
				goalPos = statePos
				grid[i][j] = "z"
				continue
			}

		}
	}

	graph := make(map[position]neighborStateData)
	for i := range grid {
		for j := range grid[i] {
			curPos := position{i, j}
			if curPos == goalPos {
				continue
			}
			statePos := position{i, j}

			neighborStateData := getNeighborStateData(statePos, grid)
			graph[statePos] = neighborStateData
		}
	}

	return state{
		grid:            grid,
		initialStatePos: initialStatePos,
		goalPos:         goalPos,
		values:          values,
		graph:           graph,
	}
}

func canBeNeighbor(stateElevation, possibleNeighborElevation string) bool {
	if stateElevation == "S" {
		return true
	}
	if possibleNeighborElevation == "E" {
		return stateElevation == "z"
	}

	stateCharCode := []rune(stateElevation)[0]
	possibleNeighborCharCode := []rune(possibleNeighborElevation)[0]

	if stateCharCode >= possibleNeighborCharCode {
		return true
	}

	if possibleNeighborCharCode-stateCharCode <= 1 {
		return true
	}

	return false
}

func getPossibleNeighbors(pos position, grid [][]string) []position {
	positions := make([]position, 0)
	i, j := pos.i, pos.j

	// left
	if j > 0 {
		positions = append(positions, position{i, j - 1})
	}
	// right
	if j < len(grid[0])-1 {
		positions = append(positions, position{i, j + 1})
	}
	// up
	if i > 0 {
		positions = append(positions, position{i - 1, j})
	}
	// down
	if i < len(grid)-1 {
		positions = append(positions, position{i + 1, j})
	}

	return positions
}

func getNeighborStateData(pos position, grid [][]string) neighborStateData {
	neighborData := make(neighborStateData)
	i, j := pos.i, pos.j

	// left
	if j > 0 && canBeNeighbor(grid[i][j], grid[i][j-1]) {
		neighborData[position{i, j - 1}] = ActionLeft
	}
	// right
	if j < len(grid[0])-1 && canBeNeighbor(grid[i][j], grid[i][j+1]) {
		neighborData[position{i, j + 1}] = ActionRight
	}
	// up
	if i > 0 && canBeNeighbor(grid[i][j], grid[i-1][j]) {
		neighborData[position{i - 1, j}] = ActionUp
	}
	// down
	if i < len(grid)-1 && canBeNeighbor(grid[i][j], grid[i+1][j]) {
		neighborData[position{i + 1, j}] = ActionDown
	}

	return neighborData
}

func printMatrix[T any](mat [][]T) {
	for i := range mat {
		for j := range mat[i] {
			fmt.Printf("%v ", mat[i][j])
		}
		fmt.Println()
	}
}

func SolveA() {
	path := "12/input/real.txt"
	lines := utils.ReadLines(path)

	state := NewState(lines)

	// reachable := findReachableFromGraph(state.initialStatePos, state.graph, nil)
	// fmt.Println("graph length before removal:", len(state.graph))
	// // remove unreachable from graph
	// for pos := range state.graph {
	// 	_, ok := reachable[pos]
	// 	if !ok {
	// 		delete(state.graph, pos)
	// 	}
	// }
	// fmt.Println("graph length after removal:", len(state.graph))

	// fmt.Println(state.grid)
	fmt.Println(state.gridString())
	// fmt.Println(state.values)
	fmt.Println(state.initialStatePos, state.goalPos)

	fmt.Println(state.grid[state.initialStatePos.i][state.initialStatePos.j], state.grid[state.goalPos.i][state.goalPos.j])

	optimalPlanGrid := state.findOptimalPlan()
	minCostInitialState := state.values[state.initialStatePos.i][state.initialStatePos.j]

	fmt.Println("values:")
	printMatrix(state.values)
	fmt.Println()

	fmt.Println("plan:")
	printMatrix(optimalPlanGrid)
	fmt.Println()

	fmt.Println("minimum cost from initial state:", minCostInitialState)

	fmt.Println("simulate plan:")
	ok, finalCost := state.simulatePlan(optimalPlanGrid, true)
	if ok {
		fmt.Println("plan is valid! Final accumulated cost:", finalCost)
	}
}
