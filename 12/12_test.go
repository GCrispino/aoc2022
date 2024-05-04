package day12

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBeNeighbor(t *testing.T) {
	type testData struct {
		stateElevation, possibleNeighborElevation string
		expectedResult                            bool
	}

	tts := []testData{
		{
			stateElevation: "S", possibleNeighborElevation: "a",
			expectedResult: true,
		},
		{
			stateElevation: "S", possibleNeighborElevation: "s",
			expectedResult: true,
		},
		{
			stateElevation: "a", possibleNeighborElevation: "b",
			expectedResult: true,
		},
		{
			stateElevation: "a", possibleNeighborElevation: "c",
			expectedResult: false,
		},
		{
			stateElevation: "a", possibleNeighborElevation: "d",
			expectedResult: false,
		},
		{
			stateElevation: "a", possibleNeighborElevation: "E",
			expectedResult: false,
		},
		{
			stateElevation: "s", possibleNeighborElevation: "E",
			expectedResult: false,
		},
		{
			stateElevation: "z", possibleNeighborElevation: "E",
			expectedResult: true,
		},
	}

	for i, tt := range tts {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := canBeNeighbor(tt.stateElevation, tt.possibleNeighborElevation)
			assert.Equal(t, tt.expectedResult, actual)
		})
	}

}

func TestGetNeighborStateData(t *testing.T) {
	grid := [][]string{
		{"S", "a", "b", "q", "p", "o", "n", "m"},
		{"a", "b", "c", "r", "y", "x", "x", "l"},
		{"a", "c", "c", "s", "z", "E", "x", "k"},
		{"a", "c", "c", "t", "u", "v", "w", "j"},
		{"a", "b", "d", "e", "f", "g", "h", "i"},
	}
	type testData struct {
		statePos                  position
		expectedNeighborStateData neighborStateData
	}

	tts := []testData{
		{
			statePos: position{0, 0},
			expectedNeighborStateData: neighborStateData{
				{0, 1}: ActionRight,
				{1, 0}: ActionDown,
			},
		},
		{
			statePos: position{0, 7},
			expectedNeighborStateData: neighborStateData{
				{0, 6}: ActionLeft,
				{1, 7}: ActionDown,
			},
		},
		{
			statePos: position{2, 3},
			expectedNeighborStateData: neighborStateData{
				{2, 2}: ActionLeft,
				{1, 3}: ActionUp,
				{3, 3}: ActionDown,
			},
		},
		{
			statePos: position{3, 3},
			expectedNeighborStateData: neighborStateData{
				{3, 2}: ActionLeft,
				{3, 4}: ActionRight,
				{2, 3}: ActionUp,
				{4, 3}: ActionDown,
			},
		},
		{
			statePos: position{4, 0},
			expectedNeighborStateData: neighborStateData{
				{4, 1}: ActionRight,
				{3, 0}: ActionUp,
			},
		},
		{
			statePos: position{4, 7},
			expectedNeighborStateData: neighborStateData{
				{4, 6}: ActionLeft,
				{3, 7}: ActionUp,
			},
		},
	}

	for i, tt := range tts {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := getNeighborStateData(tt.statePos, grid)
			assert.Equal(t, tt.expectedNeighborStateData, actual)
		})
	}

}

func TestNewStateGrid(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	lines := strings.Split(input, "\n")

	state := NewState(lines)

	expectedGrid := [][]string{
		{"a", "a", "b", "q", "p", "o", "n", "m"},
		{"a", "b", "c", "r", "y", "x", "x", "l"},
		{"a", "c", "c", "s", "z", "z", "x", "k"},
		{"a", "c", "c", "t", "u", "v", "w", "j"},
		{"a", "b", "d", "e", "f", "g", "h", "i"},
	}
	expectedGoalPos := position{2, 5}
	expectedInitialStatePos := position{0, 0}
	assert.Equal(t, expectedGrid, state.grid)
	assert.Equal(t, expectedGoalPos, state.goalPos)
	assert.Equal(t, expectedInitialStatePos, state.initialStatePos)

}

func TestFindReachable(t *testing.T) {
	input := `Sab
abc
ddE`
	lines := strings.Split(input, "\n")

	state := NewState(lines)
	expectedVisited := map[position]struct{}{
		{0, 0}: {},
		{0, 1}: {},
		{0, 2}: {},
		{1, 0}: {},
		{1, 1}: {},
		{1, 2}: {},
	}
	actualVisited := findReachableFromGraph(state.initialStatePos, state.graph, nil)
	fmt.Println("expected visited:")
	fmt.Println(expectedVisited)
	fmt.Println()
	fmt.Println("actual visited:")
	fmt.Println(actualVisited)
	fmt.Println()

	fmt.Println(state.graph)
	assert.Equal(t, expectedVisited, actualVisited)
}

func TestNewState(t *testing.T) {
	input := `Sab
abc
acE`
	lines := strings.Split(input, "\n")

	state := NewState(lines)

	expectedGrid := [][]string{
		{"a", "a", "b"},
		{"a", "b", "c"},
		{"a", "c", "z"},
	}
	expectedGraph := map[position]neighborStateData{
		{0, 0}: {
			{0, 1}: ActionRight,
			{1, 0}: ActionDown,
		},
		{0, 1}: {
			{0, 0}: ActionLeft,
			{0, 2}: ActionRight,
			{1, 1}: ActionDown,
		},
		{0, 2}: {
			{0, 1}: ActionLeft,
			{1, 2}: ActionDown,
		},

		{1, 0}: {
			{0, 0}: ActionUp,
			{1, 1}: ActionRight,
			{2, 0}: ActionDown,
		},
		{1, 1}: {
			{0, 1}: ActionUp,
			{1, 0}: ActionLeft,
			{1, 2}: ActionRight,
			{2, 1}: ActionDown,
		},
		{1, 2}: {
			{0, 2}: ActionUp,
			{1, 1}: ActionLeft,
		},

		{2, 0}: {
			{1, 0}: ActionUp,
		},
		{2, 1}: {
			{2, 0}: ActionLeft,
			{1, 1}: ActionUp,
		},
	}
	expectedGoalPos := position{2, 2}
	expectedInitialStatePos := position{0, 0}

	assert.Equal(t, expectedGrid, state.grid)
	assert.Equal(t, expectedGoalPos, state.goalPos)
	assert.Equal(t, expectedInitialStatePos, state.initialStatePos)

	for i := range state.grid {
		for j := range state.grid[i] {
			pos := position{i, j}
			assert.Equal(t, expectedGraph[pos], state.graph[pos])
		}
	}
}

func TestFindOptimalPlan(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	lines := strings.Split(input, "\n")

	state := NewState(lines)

	expectedOptimalPlanGrid := [][]action{
		{"v", ".", ".", "v", "<", "<", "<", "<"},
		{">", "v", ".", "v", "v", "<", "<", "^"},
		{".", ">", "v", "v", ">", "E", "^", "^"},
		{".", ".", "v", ">", ">", ">", "^", "^"},
		{".", ".", ">", ">", ">", ">", ">", "^"},
	}
	actualOptimalPlanGrid := state.findOptimalPlan()
	_ = actualOptimalPlanGrid
	minCostInitialState := state.values[state.initialStatePos.i][state.initialStatePos.j]

	fmt.Println("expected plan:")
	printMatrix(expectedOptimalPlanGrid)

	fmt.Println("actual plan:")
	printMatrix(actualOptimalPlanGrid)

	// assert.Equal(t, expectedOptimalPlanGrid, actualOptimalPlanGrid)
	assert.Equal(t, 31, minCostInitialState)

}

func TestGetMinimumCostLowestElevation(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	lines := strings.Split(input, "\n")

	state := NewState(lines)

	state.findOptimalPlan()
	minCostInitialState := state.values[state.initialStatePos.i][state.initialStatePos.j]
	minCostLowestElevation := getMinimumCostLowestElevation(state.grid, state.values)

	// assert.Equal(t, expectedOptimalPlanGrid, actualOptimalPlanGrid)
	assert.Equal(t, 31, minCostInitialState)
	assert.Equal(t, 29, minCostLowestElevation)

}
