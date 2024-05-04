package main

import (
	day1 "aoc-2022/1"
	day10 "aoc-2022/10"
	day11 "aoc-2022/11"
	day12 "aoc-2022/12"
	day2 "aoc-2022/2"
	day3 "aoc-2022/3"
	day4 "aoc-2022/4"
	day5 "aoc-2022/5"
	day6 "aoc-2022/6"
	day7 "aoc-2022/7"
	day8 "aoc-2022/8"
	day9 "aoc-2022/9"

	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("USAGE: ./aoc-2022 <day> <challenge>")
		return
	}

	day := args[1]

	challenge := "a"
	if len(args) == 3 {
		challenge = args[2]
	}
	_ = challenge

	switch day {
	case "1":
		if strings.ToUpper(challenge) == "B" {
			day1.SolveB()
			break
		}
		day1.SolveA()
	case "2":
		if strings.ToUpper(challenge) == "B" {
			day2.SolveB()
			break
		}
		day2.SolveA()
	case "3":
		if strings.ToUpper(challenge) == "B" {
			day3.SolveB()
			break
		}
		day3.SolveA()
	case "4":
		if strings.ToUpper(challenge) == "B" {
			day4.SolveB()
			break
		}
		day4.SolveA()
	case "5":
		if strings.ToUpper(challenge) == "B" {
			day5.SolveB()
			break
		}
		day5.SolveA()
	case "6":
		if strings.ToUpper(challenge) == "B" {
			day6.SolveB()
			break
		}
		day6.SolveA()
	case "7":
		if strings.ToUpper(challenge) == "B" {
			day7.SolveB()
			break
		}
		day7.SolveA()
	case "8":
		if strings.ToUpper(challenge) == "B" {
			day8.SolveB()
			break
		}
		day8.SolveA()
	case "9":
		if strings.ToUpper(challenge) == "B" {
			day9.SolveB()
			break
		}
		day9.SolveA()
	case "10":
		if strings.ToUpper(challenge) == "B" {
			day10.SolveB()
			break
		}
		day10.SolveA()
	case "11":
		if strings.ToUpper(challenge) == "B" {
			fmt.Println("day 11 does not have a solution implementation for challenge b yet :(")
			// day11.SolveB()
			break
		}
		day11.SolveA()
	case "12":
		if strings.ToUpper(challenge) == "B" {
			day12.SolveB()
			break
		}
		day12.SolveA()
	default:
		fmt.Printf("day %s does not have a solution implementation!", day)
		os.Exit(1)
	}

}
