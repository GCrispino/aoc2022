package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	// "strings"

	"aoc-2022/pkg/utils"
)

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

func (p position) move(d direction) position {
	newPos := p
	switch d {
	case RightDirection:
		newPos.col += 1
	case LeftDirection:
		newPos.col -= 1
	case UpDirection:
		newPos.row += 1
	case DownDirection:
		newPos.row -= 1
	default:
		panic("invalid direction")
	}

	return newPos
}

type ropeState struct {
	positions      []position
	visitedTracker posTracker
}

func NewRopeState(nKnots int) ropeState {
	return ropeState{
		positions:      make([]position, nKnots),
		visitedTracker: make(posTracker),
	}
}

func (r *ropeState) Transition(m move) {
	// fmt.Printf("initial position: %v, move: %v\n", r, m)
	for i := 0; i < m.n; i++ {
		headPosition := r.positions[0]
		// move head's position
		r.positions[0] = headPosition.move(m.dir)
		for j := 0; j < len(r.positions)-1; j++ {

			// update tail's position
			r.positions[j+1] = r.getNewTailPosition(j)
			// fmt.Printf("position at %d: %v\n", i, r)

			if j == len(r.positions)-2 {
				tailPos := r.positions[j+1]
				_, ok := r.visitedTracker[tailPos]
				if !ok {
					r.visitedTracker[tailPos] = struct{}{}
				}
			}
		}
	}
	// fmt.Println()
}

func (r *ropeState) getNewTailPosition(i int) position {
	// fmt.Println("rowState: ", r)
	headPos := r.positions[i]
	curTailPos := r.positions[i+1]

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

func simulate(moves []move, nKnots int) ropeState {
	state := NewRopeState(nKnots)

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

	finalState := simulate(moves, 2)

	fmt.Println(finalState.visitedTracker.String())

	fmt.Println(len(finalState.visitedTracker))
}
