package main

import (
	"aoc-2022/pkg/utils"
)

func findCommonItemStrings(strs []string) rune {
	if len(strs) < 2 {
		panic("slice lenght should be higher than 1")
	}

	setMaps := make([]map[rune]struct{}, len(strs))

	for i, s := range strs {
		setMaps[i] = utils.BuildSetMapFromContainer([]rune(s))
	}

	// var common rune
	commons := make([]rune, 0)
	for i := range setMaps {
		if i == len(setMaps)-1 {
			break
		}

		firstMap := setMaps[i]
		secondMap := setMaps[i+1]
		for item := range firstMap {
			if _, ok := secondMap[item]; ok {
				commons = append(commons, item)
			}
		}
	}

	return commons[0]
}
