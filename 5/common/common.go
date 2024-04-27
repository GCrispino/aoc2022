package common

import (
	"regexp"
	"strconv"
	"strings"

	"aoc-2022/pkg/utils/types"
)

type problemData struct {
	stacks []crateStack
	plan   rearrangingPlan
}

type strategy int

const (
	S9000 = iota
	S9001
)

func (p problemData) ApplyPlan(strat strategy) {
	for _, step := range p.plan.steps {
		nCratesToMove := step.cratesToMove
		fromStack := &p.stacks[step.from-1]
		toStack := &p.stacks[step.to-1]

		if strat == S9000 {
			for i := 0; i < nCratesToMove; i++ {
				val := fromStack.Pop()
				toStack.Push(val)
			}
		} else if strat == S9001 {
			newS := types.NewStack[string]()
			for i := 0; i < nCratesToMove; i++ {
				val := fromStack.Pop()
				newS.Push(val)
			}

			for i := 0; i < nCratesToMove; i++ {
				val := newS.Pop()
				toStack.Push(val)
			}
		}

	}
}

func (p problemData) GetTop() string {
	topCrates := make([]string, len(p.stacks))
	for i, stack := range p.stacks {
		topCrates[i] = stack.Peek()
	}

	return strings.Join(topCrates, "")
}

type crateStack struct {
	id int
	types.Stack[string]
}

type step struct {
	cratesToMove, from, to int
}

type rearrangingPlan struct {
	steps []step
}

func parsePlanFromLines(planLines []string) rearrangingPlan {
	r := rearrangingPlan{
		steps: make([]step, len(planLines)),
	}
	for i, line := range planLines {
		spl := strings.Split(line, " ")

		cratesToMove, err := strconv.Atoi(spl[1])
		if err != nil {
			panic(err)
		}

		from, err := strconv.Atoi(spl[3])
		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(spl[5])
		if err != nil {
			panic(err)
		}

		r.steps[i] = step{
			cratesToMove: cratesToMove,
			from:         from,
			to:           to,
		}
	}

	return r
}

func parseStacksFromLines(lines []string, iSepLine int) []crateStack {
	stackNumbersLine := lines[iSepLine-1]
	iLastStackNumber := len(stackNumbersLine) - 2
	nStacks, err := strconv.Atoi(stackNumbersLine[iLastStackNumber : iLastStackNumber+1])
	if err != nil {
		panic(err)
	}

	numberPattern := regexp.MustCompile(`\d`)

	matches := numberPattern.FindAllStringIndex(stackNumbersLine, -1)

	stacks := make([]crateStack, nStacks)
	for i, match := range matches {
		n, err := strconv.Atoi(stackNumbersLine[match[0]:match[1]])
		if err != nil {
			panic(err)
		}

		stacks[i] = crateStack{
			id:    n,
			Stack: types.NewStack[string](),
		}
		for iLine := iSepLine - 2; iLine >= 0; iLine-- {
			if crate := string(lines[iLine][match[0]]); crate != " " {
				stacks[i].Push(string(crate))
			}
		}
	}

	return stacks
}

func GetProblemData(lines []string) problemData {

	var planLines []string
	var iSepLine int
	for i, line := range lines {
		if line == "" {
			iSepLine = i
			planLines = lines[i+1:]
			break
		}
	}

	stacks := parseStacksFromLines(lines, iSepLine)
	plan := parsePlanFromLines(planLines)

	return problemData{
		stacks: stacks,
		plan:   plan,
	}
}
