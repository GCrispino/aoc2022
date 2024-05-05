package common

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	// big "github.com/ncw/gmp"

	"aoc-2022/pkg/utils"
)

type Monkey struct {
	ItemCount int

	items     []*big.Int
	operation ItemOp
	test      ItemTest
}

func NewMonkeyFromLines(lines []string, i int) Monkey {
	itemsStr := strings.Split(
		strings.Split(lines[i+1], ": ")[1],
		", ",
	)

	// items, err := utils.MapWithError(itemsStr, strconv.Atoi)
	items, err := utils.MapWithError(itemsStr, func(s string) (*big.Int, error) {
		x, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		return big.NewInt(int64(x)), nil
	})
	utils.CheckAndPanic(err)

	return Monkey{
		items:     items,
		operation: NewItemOp(lines[i+2]),
		test:      NewItemTest(lines[i+3 : i+6]),
	}
}

type ItemOp func(*big.Int) *big.Int

func NewItemOp(opDesc string) ItemOp {
	spl := strings.Split(opDesc, " = ")
	spl2 := strings.Split(spl[1], " ")

	operand := spl2[2]

	operator := spl2[1]

	return func(old *big.Int) *big.Int {
		var opBInt *big.Int
		var isOld bool
		if operand == "old" {
			opBInt = old
			isOld = true
		} else {
			opInt, err := strconv.Atoi(operand)
			if err != nil {
				panic(err)
			}

			opBInt = big.NewInt(int64(opInt))
		}

		// fmt.Println("  itemop", old, operand)

		switch operator {
		case "+":
			if isOld {
				return old.Lsh(old, 1)
			}

			return old.Add(old, opBInt)
		case "*":
			if isOld {
				return old.Exp(old, big.NewInt(2), nil)
			}

			return old.Mul(old, opBInt)
		default:
			panic("invalid operator")
		}
	}
}

type ItemTest struct {
	divisor                         int
	ifTrueMonkeyId, ifFalseMonkeyId int
}

func NewItemTest(itemTestDesc []string) ItemTest {
	divisor := utils.SplitAndGetNthItemInt(itemTestDesc[0], "divisible by ")

	ifTrueMonkeyId := utils.SplitAndGetNthItemInt(itemTestDesc[1], "throw to monkey ")

	ifFalseMonkeyId := utils.SplitAndGetNthItemInt(itemTestDesc[2], "throw to monkey ")

	return ItemTest{
		divisor:         divisor,
		ifTrueMonkeyId:  ifTrueMonkeyId,
		ifFalseMonkeyId: ifFalseMonkeyId,
	}
}

const nLinesMonkeyInput int = 6

func ParseInput(lines []string) []*Monkey {
	nLines := len(lines)
	monkeysData := make([]*Monkey, 0)

	for i := 0; i < nLines; i += nLinesMonkeyInput + 1 {
		m := NewMonkeyFromLines(lines, i)
		monkeysData = append(monkeysData, &m)
	}

	return monkeysData
}

func InvestigateMonkeys(monkeysData []*Monkey, worryLevelDivisor, nRounds int) []*Monkey {
	rateToPrint := 1
	for i := 0; i < nRounds; i++ {
		if i%rateToPrint == 0 {
			fmt.Println("round", i, monkeysData[0].items, monkeysData[1].items)
		}
		for _, m := range monkeysData {
			// fmt.Printf("monkey %d. Items: %v\n", iMonkey, m.items)
			for _, item := range m.items {
				m.ItemCount++

				// update worry level
				// fmt.Println("item:", item)
				newWorryLevel := m.operation(item)
				if worryLevelDivisor != 1 {
					newWorryLevel.Div(newWorryLevel, big.NewInt(int64(worryLevelDivisor)))
				}
				// newWorryLevel := m.operation(item) / worryLevelDivisor
				// fmt.Println(" new worry level:", item, newWorryLevel)

				// test for which monkey to throw item
				var toThrow int
				mod := new(big.Int).Mod(newWorryLevel, big.NewInt(int64(m.test.divisor)))
				if mod.Cmp(big.NewInt(0)) == 0 {
					// fmt.Println(newWorryLevel, m.test.divisor)
					toThrow = m.test.ifTrueMonkeyId
				} else {
					toThrow = m.test.ifFalseMonkeyId
				}

				// apply throw
				monkeysData[toThrow].items = append(monkeysData[toThrow].items, newWorryLevel)
				// fmt.Println("oi", m.items, iItem, item, newWorryLevel)
			}

			// clear items that were throwed to other monkeys
			m.items = make([]*big.Int, 0)
			// fmt.Printf("monkey %d. Items: %v\n", iMonkey, m.items)
			// for jMonkey := range monkeysData {
			// 	fmt.Printf("monkey %d. Items: %v\n", jMonkey, monkeysData[jMonkey].items)
			// }
			// fmt.Println()

			// if iMonkey == 1 {
			// 	panic("oshd")
			// }

		}
		// if i == 3 {
		// 	break
		// }
	}

	return monkeysData
}
