package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	// "strings"

	"aoc-2022/pkg/utils"
)

/*
 * TODO
 *
 * Write transition function
 * - receives direction (R, L, U, D) and number
 * -  computes new position of head
 * -  computes new position of tail
 *	- this can probably be its own function
 * Write simulation function
 * - goes through each command and computes the transition for each
 * - needs to keep track of the tail position for computing the final answer
 */

type direction string

const (
	RightDirection direction = "R"
	LeftDirection  direction = "L"
	UpDirection    direction = "U"
	DownDirection  direction = "D"
)

type move struct {
	dir direction
	n   int
}

type position struct {
	row, col int
}

func (p *position) move(d direction) {
	switch d {
	case RightDirection:
		p.col += 1
	case LeftDirection:
		p.col -= 1
	case UpDirection:
		p.row += 1
	case DownDirection:
		p.row -= 1
	default:
		panic("invalid direction")
	}
}

type ropeState struct {
	headPosition, tailPosition position
	visitedTracker             posTracker
}

func NewRopeState(headPos, tailPos position) ropeState {
	return ropeState{
		headPosition:   headPos,
		tailPosition:   tailPos,
		visitedTracker: make(posTracker),
	}
}

func (r *ropeState) Transition(m move) {
	// fmt.Printf("initial position: %v, move: %v\n", r, m)
	for i := 0; i < m.n; i++ {
		// move head's position
		r.headPosition.move(m.dir)

		// update tail's position
		r.tailPosition = r.getNewTailPosition()
		// fmt.Printf("position at %d: %v\n", i, r)

		tailPos := r.tailPosition
		_, ok := r.visitedTracker[tailPos]
		if !ok {
			r.visitedTracker[tailPos] = struct{}{}
		}
	}
	// fmt.Println()
}

func (r *ropeState) getNewTailPosition() position {
	// fmt.Println("rowState: ", r)
	headPos := r.headPosition
	curTailPos := r.tailPosition

	newTailPos := curTailPos

	// get diffs
	colDiff := headPos.col - curTailPos.col
	rowDiff := headPos.row - curTailPos.row

	absColDiff := math.Abs(float64(colDiff))
	absRowDiff := math.Abs(float64(rowDiff))
	// fmt.Printf("  rowDiff: %d, colDiff: %d\n", rowDiff, colDiff)

	if colDiff > 1 || colDiff > 0 && absRowDiff > 1 {
		newTailPos.col += 1
	} else if colDiff < -1 || colDiff < 0 && absRowDiff > 1 {
		newTailPos.col -= 1
	}

	if rowDiff > 1 || rowDiff > 0 && absColDiff > 1 {
		newTailPos.row += 1
	} else if rowDiff < -1 || rowDiff < 0 && absColDiff > 1 {
		newTailPos.row -= 1
	}

	// fmt.Println()
	return newTailPos
}

type posTracker map[position]struct{}

func (t posTracker) String() string {
	maxCol, maxRow := 0, 0
	// fmt.Println("tracked positions!")
	for pos := range t {
		// fmt.Println("  ", pos)
		row, col := pos.row, pos.col
		if row > maxRow {
			maxRow = row
		}

		if col > maxCol {
			maxCol = col
		}
	}
	// fmt.Println()

	matrixStr := ""
	// matrix := make([][]string, maxRow+1)
	for i := maxRow; i >= 0; i-- {
		// matrixStr += strings.Join(matrix[i], "") + "\n"
		matrixRowStr := ""
		for j := 0; j <= maxCol+1; j++ {
			_, ok := t[position{i, j}]
			if ok {
				matrixRowStr += "#"
			} else {
				matrixRowStr += "."
			}
		}

		if i == 0 {
			matrixRowStr = "s" + matrixRowStr[1:]
			matrixStr += matrixRowStr
		} else {
			matrixStr += matrixRowStr + "\n"
		}
	}

	return matrixStr
}

func (t posTracker) getCount() (count int) {
	for range t {
		count += 1
	}

	return count
}

func simulate(moves []move) ropeState {
	state := NewRopeState(position{}, position{})

	for _, m := range moves {
		state.Transition(m)
	}

	return state
}

func parseMovesFromFile(path string) []move {
	lines := utils.ReadLines(path)
	moves := make([]move, len(lines))
	for i, line := range lines {
		spl := strings.Split(line, " ")

		dir := direction(spl[0])
		switch dir {
		case RightDirection:
		case LeftDirection:
		case UpDirection:
		case DownDirection:
		default:
			panic(fmt.Errorf("direction %s not supported", dir))
		}

		n, err := strconv.Atoi(spl[1])
		if err != nil {
			panic(err)
		}

		moves[i] = move{
			dir: dir,
			n:   n,
		}
	}

	return moves
}

func SolveA() {
	moves := parseMovesFromFile("9/input/a.txt")

	finalState := simulate(moves)

	fmt.Println(finalState.visitedTracker.String())

	fmt.Println(len(finalState.visitedTracker))
}
