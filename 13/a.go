package day13

import (
	"encoding/json"
	"fmt"

	"aoc-2022/pkg/utils"
)

type rawPacketData []any

type packetData struct {
	v *int
	l []packetData
}

func (p packetData) String() string {
	if p.v != nil {
		return fmt.Sprintf("%d", *p.v)
	}

	if p.l != nil {
		s := "["
		for i, data := range p.l {
			s += data.String()
			if i != len(p.l)-1 {
				s += ", "
			}
		}
		s += "]"

		return s
	}

	return ""
}

type packetDataPair struct {
	packetData1, packetData2 packetData
}

func (p packetDataPair) String() string {
	return fmt.Sprintf("%s\n%s\n", p.packetData1, p.packetData2)
}

type order int

const (
	orderCorrect order = iota
	orderIncorrect
	orderContinue
)

func (p *packetDataPair) checkOrder() order {
	return checkOrder(p.packetData1, p.packetData2)
}

func checkOrder(packetData1, packetData2 packetData) order {
	is1Number := packetData1.v != nil
	is1List := packetData1.l != nil

	is2Number := packetData2.v != nil
	is2List := packetData2.l != nil

	// fmt.Println(packetData1.v, packetData1.l, packetData2.v, packetData2.l)
	// fmt.Printf("%v, %v, %t, %t, %t, %t\n", packetData1, packetData2, is1Number, is1List, is2Number, is2List)

	switch {
	case is1Number && is2Number:
		if *packetData1.v == *packetData2.v {
			return orderContinue
		}
		if *packetData1.v > *packetData2.v {
			return orderIncorrect
		}

		return orderCorrect
	case is1List && is2Number:
		list2 := packetData{l: []packetData{
			{
				v: packetData2.v,
			},
		}}
		return checkOrder(packetData1, list2)
	case is2List && is1Number:
		list1 := packetData{l: []packetData{
			{
				v: packetData1.v,
			},
		}}
		return checkOrder(list1, packetData2)
	case is1List && is2List:
		l1 := packetData1.l
		l2 := packetData2.l
		i := 0
		for {
			// fmt.Printf("  %d, %v, %v, %d, %d", i, l1, l2, len(l1), len(l2))
			if i == len(l1) && i == len(l2) {
				return orderContinue
			}
			if i == len(l1) && i < len(l2) {
				return orderCorrect
			}
			if i == len(l2) && i < len(l1) {
				return orderIncorrect
			}
			orderCurrentItems := checkOrder(l1[i], l2[i])
			if orderCurrentItems == orderContinue {
				i++
				continue
			}

			return orderCurrentItems
		}
	}

	panic("shouldn't happen")
}

func newPacketDataFromAny(data any) packetData {
	v := packetData{}
	switch d := data.(type) {
	case []any:
		v.l = NewPacketDataFromRaw(d).l
	case float64:
		v.v = new(int)
		*v.v = int(d)
	default:
		panic("unsupported")
	}

	return v
}

func NewPacketDataFromRaw(data rawPacketData) packetData {
	return packetData{
		l: utils.Map(data, newPacketDataFromAny),
	}
}

func NewPacketDataFromString(s string) packetData {
	var r rawPacketData
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		panic(err)
	}
	return NewPacketDataFromRaw(r)
}

func pairFromLines(line1, line2 string) packetDataPair {
	return packetDataPair{
		NewPacketDataFromString(line1),
		NewPacketDataFromString(line2),
	}
}

func getInput(filePath string) []packetDataPair {
	lines := utils.ReadLines(filePath)
	rawPacketVals := make([]packetDataPair, 0)
	for i := 0; i < len(lines); i += 3 {
		rawPacketVals = append(rawPacketVals, pairFromLines(lines[i], lines[i+1]))
	}

	return rawPacketVals
}

func checkOrderPairs(pairs []packetDataPair) []int {
	indices := make([]int, 0)

	for i, pair := range pairs {
		if pair.checkOrder() == orderCorrect {
			indices = append(indices, i+1)
		}
	}

	return indices
}

func SolveA() {
	packetPairs := getInput("13/input/input.txt")

	indices := checkOrderPairs(packetPairs)
	fmt.Println("indices of pairs with correct order:", indices)

	sum := 0
	for _, i := range indices {
		sum += i
	}
	fmt.Println("sum: ", sum)
}
