package day9

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTailPosition(t *testing.T) {
	type testData struct {
		name                    string
		curState, expectedState ropeState
	}

	tts := []testData{
		{
			/**
			 * .....    .....    .....
			 * .TH.. -> .T.H. -> ..TH.
			 * .....    .....    .....
			 */
			name: "updates tail position correctly - 1",
			curState: ropeState{
				positions: []position{
					{row: 1, col: 3},
					{row: 1, col: 1},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 1, col: 3},
					{row: 1, col: 2},
				},
			},
		},
		{
			/**
			 * ...    ...    ...
			 * .T.    .T.    ...
			 * .H. -> ... -> .T.
			 * ...    .H.    .H.
			 * ...    ...    ...
			 */
			name: "updates tail position correctly - 2",
			curState: ropeState{
				positions: []position{
					{row: 1, col: 1},
					{row: 3, col: 1},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 1, col: 1},
					{row: 2, col: 1},
				},
			},
		},
		{
			/**
			 * .....    .....    .....
			 * .....    ..H..    ..H..
			 * ..H.. -> ..... -> ..T..
			 * .T...    .T...    .....
			 * .....    .....    .....
			 */
			name: "updates tail position correctly - 3",
			curState: ropeState{
				positions: []position{
					{row: 3, col: 2},
					{row: 1, col: 1},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 3, col: 2},
					{row: 2, col: 2},
				},
			},
		},
		{
			/**
			 * .....    .....    .....
			 * .....    .....    .....
			 * ..H.. -> ...H. -> ..TH.
			 * .T...    .T...    .....
			 * .....    .....    .....
			 */
			name: "updates tail position correctly - 4",
			curState: ropeState{
				positions: []position{
					{row: 2, col: 3},
					{row: 1, col: 1},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 2, col: 3},
					{row: 2, col: 2},
				},
			},
		},
		{
			/**
			 * ...H..   ..H...   ..HT..
			 * ....T.   ....T.   ......
			 * ...... ->...... ->......
			 * ......   ......   ......
			 * ......   ......   ......
			 */
			name: "updates tail position correctly - 5",
			curState: ropeState{
				positions: []position{
					{row: 4, col: 2},
					{row: 3, col: 4},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 4, col: 2},
					{row: 4, col: 3},
				},
			},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			state := tt.curState
			expected := tt.expectedState
			actual := ropeState{
				positions: []position{
					state.positions[0],
					state.getNewTailPosition(0),
				},
			}

			assert.Equal(t, expected.positions[0], actual.positions[0])
			assert.Equal(t, expected.positions[1], actual.positions[1])
		})
	}

}

func TestTransition(t *testing.T) {
	type testData struct {
		name                    string
		move                    move
		curState, expectedState ropeState
	}

	tts := []testData{
		{
			name: "moves and updates tail position correctly - 1",
			curState: ropeState{
				positions: []position{
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
			move: move{dir: RightDirection, n: 4},
			expectedState: ropeState{
				positions: []position{
					{row: 0, col: 4},
					{row: 0, col: 3},
				},
			},
		},
		{
			name: "moves and updates tail position correctly - 2",
			curState: ropeState{
				positions: []position{
					{row: 0, col: 4},
					{row: 0, col: 3},
				},
			},
			move: move{dir: UpDirection, n: 4},
			expectedState: ropeState{
				positions: []position{
					{row: 4, col: 4},
					{row: 3, col: 4},
				},
			},
		},
		{
			name: "moves and updates tail position correctly - 3",
			curState: ropeState{
				positions: []position{
					{row: 4, col: 4},
					{row: 3, col: 4},
				},
			},
			move: move{dir: LeftDirection, n: 3},
			expectedState: ropeState{
				positions: []position{
					{row: 4, col: 1},
					{row: 4, col: 2},
				},
			},
		},
		{
			name: "moves and updates tail position correctly - 4",
			curState: ropeState{
				positions: []position{
					{row: 4, col: 1},
					{row: 4, col: 2},
				},
			},
			move: move{dir: DownDirection, n: 1},
			expectedState: ropeState{
				positions: []position{
					{row: 3, col: 1},
					{row: 4, col: 2},
				},
			},
		},
		{
			name:     "moves and updates tail position correctly with more than 2 knots - 1 - just one move",
			curState: NewRopeState(10),
			move:     move{dir: RightDirection, n: 1},
			expectedState: ropeState{
				positions: []position{
					{row: 0, col: 1},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
		},
		{
			name:     "moves and updates tail position correctly with more than 2 knots - 1 - four moves",
			curState: NewRopeState(10),
			move:     move{dir: RightDirection, n: 4},
			expectedState: ropeState{
				positions: []position{
					{row: 0, col: 4},
					{row: 0, col: 3},
					{row: 0, col: 2},
					{row: 0, col: 1},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
		},
		{
			name: "moves and updates tail position correctly with more than 2 knots - 2",
			move: move{dir: UpDirection, n: 4},
			curState: ropeState{
				positions: []position{
					{row: 0, col: 4},
					{row: 0, col: 3},
					{row: 0, col: 2},
					{row: 0, col: 1},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 4, col: 4},
					{row: 3, col: 4},
					{row: 2, col: 4},
					{row: 2, col: 3},
					{row: 2, col: 2},
					{row: 1, col: 1},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
		},
		{
			name: "moves and updates tail position correctly with more than 2 knots - 3",
			move: move{dir: LeftDirection, n: 3},
			curState: ropeState{
				positions: []position{
					{row: 4, col: 4},
					{row: 3, col: 4},
					{row: 2, col: 4},
					{row: 2, col: 3},
					{row: 2, col: 2},
					{row: 1, col: 1},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
			expectedState: ropeState{
				positions: []position{
					{row: 4, col: 1},
					{row: 4, col: 2},
					{row: 3, col: 3},
					{row: 2, col: 3},
					{row: 2, col: 2},
					{row: 1, col: 1},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
					{row: 0, col: 0},
				},
			},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			state := tt.curState
			state.visitedTracker = make(posTracker)
			expected := tt.expectedState

			state.Transition(tt.move)
			actual := state

			for i, actualPos := range actual.positions {
				assert.Equal(t, expected.positions[i], actualPos, "i = %d", i)
			}
		})
	}

}

func TestSimulate(t *testing.T) {
	moves := []move{
		{dir: RightDirection, n: 4},
		{dir: UpDirection, n: 4},
		{dir: LeftDirection, n: 3},
		{dir: DownDirection, n: 1},
		{dir: RightDirection, n: 4},
		{dir: DownDirection, n: 1},
		{dir: LeftDirection, n: 5},
		{dir: RightDirection, n: 2},
	}

	expectedState := ropeState{
		positions: []position{
			{row: 2, col: 2},
			{row: 2, col: 1},
		},
	}
	newState := simulate(moves, 2)
	actualPosTracker := newState.visitedTracker

	// fmt.Printf("tracker! %v\n", actualPosTracker)
	expectedPosTrackerCount := 13
	assert.Equal(t, expectedPosTrackerCount, actualPosTracker.getCount())

	assert.Equal(t, expectedState.positions[0], newState.positions[0])
	assert.Equal(t, expectedState.positions[1], newState.positions[1])

	expectedPosTrackerStr := `..##..
...##.
.####.
....#.
s###..`

	assert.Equal(t, expectedPosTrackerStr, actualPosTracker.String())

}

func TestSimulateMultipleKnots(t *testing.T) {
	// t.Skip()
	moves := []move{
		{dir: RightDirection, n: 5},
		{dir: UpDirection, n: 8},
		{dir: LeftDirection, n: 8},
		{dir: DownDirection, n: 3},
		{dir: RightDirection, n: 17},
		{dir: DownDirection, n: 10},
		{dir: LeftDirection, n: 25},
		{dir: UpDirection, n: 20},
	}

	expectedState := ropeState{
		positions: []position{
			{row: 15, col: -11},
			{row: 14, col: -11},
			{row: 13, col: -11},
			{row: 12, col: -11},
			{row: 11, col: -11},
			{row: 10, col: -11},
			{row: 9, col: -11},
			{row: 8, col: -11},
			{row: 7, col: -11},
			{row: 6, col: -11},
		},
	}
	newState := simulate(moves, 10)
	actualPosTracker := newState.visitedTracker

	expectedPosTrackerCount := 36
	assert.Equal(t, expectedPosTrackerCount, actualPosTracker.getCount())

	for i, actualPos := range newState.positions {
		assert.Equal(t, expectedState.positions[i], actualPos, "i = %d", i)
	}

}
