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
				headPosition: position{row: 1, col: 3},
				tailPosition: position{row: 1, col: 1},
			},
			expectedState: ropeState{
				headPosition: position{row: 1, col: 3},
				tailPosition: position{row: 1, col: 2},
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
				headPosition: position{row: 1, col: 1},
				tailPosition: position{row: 3, col: 1},
			},
			expectedState: ropeState{
				headPosition: position{row: 1, col: 1},
				tailPosition: position{row: 2, col: 1},
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
				headPosition: position{row: 3, col: 2},
				tailPosition: position{row: 1, col: 1},
			},
			expectedState: ropeState{
				headPosition: position{row: 3, col: 2},
				tailPosition: position{row: 2, col: 2},
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
				headPosition: position{row: 2, col: 3},
				tailPosition: position{row: 1, col: 1},
			},
			expectedState: ropeState{
				headPosition: position{row: 2, col: 3},
				tailPosition: position{row: 2, col: 2},
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
				headPosition: position{row: 4, col: 2},
				tailPosition: position{row: 3, col: 4},
			},
			expectedState: ropeState{
				headPosition: position{row: 4, col: 2},
				tailPosition: position{row: 4, col: 3},
			},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			state := tt.curState
			expected := tt.expectedState
			actual := ropeState{
				headPosition: state.headPosition,
				tailPosition: state.getNewTailPosition(),
			}

			assert.Equal(t, expected.headPosition, actual.headPosition)
			assert.Equal(t, expected.tailPosition, actual.tailPosition)
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
				headPosition: position{row: 0, col: 0},
				tailPosition: position{row: 0, col: 0},
			},
			move: move{dir: RightDirection, n: 4},
			expectedState: ropeState{
				headPosition: position{row: 0, col: 4},
				tailPosition: position{row: 0, col: 3},
			},
		},
		{
			name: "moves and updates tail position correctly - 2",
			curState: ropeState{
				headPosition: position{row: 0, col: 4},
				tailPosition: position{row: 0, col: 3},
			},
			move: move{dir: UpDirection, n: 4},
			expectedState: ropeState{
				headPosition: position{row: 4, col: 4},
				tailPosition: position{row: 3, col: 4},
			},
		},
		{
			name: "moves and updates tail position correctly - 3",
			curState: ropeState{
				headPosition: position{row: 4, col: 4},
				tailPosition: position{row: 3, col: 4},
			},
			move: move{dir: LeftDirection, n: 3},
			expectedState: ropeState{
				headPosition: position{row: 4, col: 1},
				tailPosition: position{row: 4, col: 2},
			},
		},
		{
			name: "moves and updates tail position correctly - 4",
			curState: ropeState{
				headPosition: position{row: 4, col: 1},
				tailPosition: position{row: 4, col: 2},
			},
			move: move{dir: DownDirection, n: 1},
			expectedState: ropeState{
				headPosition: position{row: 3, col: 1},
				tailPosition: position{row: 4, col: 2},
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

			assert.Equal(t, expected.headPosition, actual.headPosition)
			assert.Equal(t, expected.tailPosition, actual.tailPosition)
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
		headPosition: position{row: 2, col: 2},
		tailPosition: position{row: 2, col: 1},
	}
	newState := simulate(moves)
	actualPosTracker := newState.visitedTracker

	// fmt.Printf("tracker! %v\n", actualPosTracker)
	expectedPosTrackerCount := 13
	assert.Equal(t, expectedPosTrackerCount, actualPosTracker.getCount())

	assert.Equal(t, expectedState.headPosition, newState.headPosition)
	assert.Equal(t, expectedState.tailPosition, newState.tailPosition)

	expectedPosTrackerStr := `..##..
...##.
.####.
....#.
s###..`

	assert.Equal(t, expectedPosTrackerStr, actualPosTracker.String())

}
