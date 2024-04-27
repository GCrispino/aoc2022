package common

import (
	"aoc-2022/pkg/utils"
)

var (
	aOffset int = int('a')
	AOffset int = int('A')
)

func FindCommonItemStrings(strs []string) rune {
	if len(strs) < 2 {
		panic("slice length should be higher than 1")
	}

	setMaps := make([]map[rune]struct{}, len(strs))

	for i, s := range strs {
		setMaps[i] = utils.BuildSetMapFromContainer([]rune(s))
	}

	// var common rune
	firstMap := make([]rune, 0)
	for item := range setMaps[0] {
		firstMap = append(firstMap, item)
	}

	secondMap := setMaps[1]

	var commons []rune
	for i := range setMaps {
		commons = make([]rune, 0)

		for _, item := range firstMap {
			if _, ok := secondMap[item]; ok {
				commons = append(commons, item)
			}
		}

		if i == len(setMaps)-2 {
			break
		}
		firstMap = commons
		secondMap = setMaps[i+2]
	}

	return commons[0]
}

func CalcPriority(val rune) (priority int) {
		priorityLower := int(val) - aOffset + 1
		priorityUpper := int(val) - AOffset + 27

		if priorityLower > 0 {
			priority = priorityLower
		} else {
			priority = priorityUpper
		}

	return
}
