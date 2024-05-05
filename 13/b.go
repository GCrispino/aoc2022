package day13

import (
	"fmt"
	"slices"
)

func SolveB() {
	packetPairs := getInput("13/input/input.txt")
	dividerPacketStr1 := "[[2]]"
	dividerPacketStr2 := "[[6]]"
	dividerPair := pairFromLines(dividerPacketStr1, dividerPacketStr2)
	mergedPacketPairs := []packetData{
		dividerPair.packetData1,
		dividerPair.packetData2,
	}
	for _, packetPair := range packetPairs {
		mergedPacketPairs = append(mergedPacketPairs, packetPair.packetData1, packetPair.packetData2)
	}

	slices.SortFunc(mergedPacketPairs, func(a, b packetData) int {
		switch checkOrder(a, b) {
		case orderCorrect:
			return -1
		case orderIncorrect:
			return 1
		case orderContinue:
			panic("should not happen")
		}

		return -1
	})

	var dividerPacket1Index, dividerPacket2Index int
	for i, packetData := range mergedPacketPairs {
		if packetData.String() == dividerPacketStr1 {
			dividerPacket1Index = i + 1
		}
		if packetData.String() == dividerPacketStr2 {
			dividerPacket2Index = i + 1
		}
	}
	fmt.Println("indices: ", dividerPacket1Index, dividerPacket2Index)
	fmt.Println("product: ", dividerPacket1Index*dividerPacket2Index)
}
