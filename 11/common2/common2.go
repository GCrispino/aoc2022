package common2

import (
	"fmt"
	"strconv"
	"strings"

	big "github.com/ncw/gmp"
	"golang.org/x/exp/slices"

	"aoc-2022/pkg/utils"
)

// type Item *big.Int

type Item struct {
	factors []*big.Int
}

func NewItem (x *big.Int) *Item {
	return &Item{factors: utils.PrimeFactors(x)}
}

func (i *Item) Add(x *big.Int) {
	e := i.Eval()

	e.Add(e, x)
	newItem := NewItem(e)
	i.factors = newItem.factors
}

func (i *Item) Add2(j *Item) {
	var kMatch int
	for k, f := range i.factors {
		if f != j.factors[k] {
			kMatch = k
			break
		}
	}
	newFactors := []*big.Int{}
	i.factors = newFactors
}

func (i *Item) Mul(x *big.Int) {

	changed := false
	for j, f := range i.factors {
		if f.Cmp(x) >= 0 {
			i.factors = append(
				i.factors[:j],
				append(
					[]*big.Int{x},
					i.factors[j:]...,
				)...,
			)

			changed = true
			break
		}
	}

	if !changed {
		i.factors = append(i.factors, x)
	}
}

func (i *Item) Exp2() {
	// TODO: possibly refactor as this is inefficient
	for _, f := range i.factors {
		i.Mul(f)
	}
}

func (i *Item) IsDivisibleByPrime(p int) bool {
	_, found := slices.BinarySearchFunc(i.factors, big.NewInt(int64(p)), func(x, y *big.Int) int {
		return x.Cmp(y)
	})

	return found
}

func (i *Item) Eval() *big.Int {
	res := big.NewInt(1)
	for _, f := range i.factors {
		res.Mul(res, f)
	}

	return res
}

type Monkey struct {
	ItemCount int

	items     []*Item
	operation ItemOp
	test      ItemTest
}

func NewMonkeyFromLines(lines []string, i int) Monkey {
	itemsStr := strings.Split(
		strings.Split(lines[i+1], ": ")[1],
		", ",
	)

	// items, err := utils.MapWithError(itemsStr, strconv.Atoi)
	items, err := utils.MapWithError(itemsStr, func(s string) (*Item, error) {
		x, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		bigX := big.NewInt(int64(x))
		factors := utils.PrimeFactors(bigX)

		return &Item{factors: factors}, nil
	})
	utils.CheckAndPanic(err)

	return Monkey{
		items:     items,
		operation: NewItemOp(lines[i+2]),
		test:      NewItemTest(lines[i+3 : i+6]),
	}
}

type ItemOp func(Item) Item

func NewItemOp(opDesc string) ItemOp {
	spl := strings.Split(opDesc, " = ")
	spl2 := strings.Split(spl[1], " ")

	operand := spl2[2]

	operator := spl2[1]

	return func(old Item) Item {
		var opBInt *big.Int
		if operand == "old" {
			old.Exp2()
			return old
		} else {
			opInt, err := strconv.Atoi(operand)
			if err != nil {
				panic(err)
			}

			opBInt = big.NewInt(int64(opInt))
		}

		switch operator {
		case "+":
			// return old.Add(old, opBInt)
			old.Add(opBInt)
		case "*":
			// return old.Mul(old, opBInt)
			old.Mul(opBInt)
		default:
			panic("invalid operator")
		}

		return old
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
		if (i+1)%rateToPrint == 0 {
			fmt.Println("round", i)
		}
		for _, m := range monkeysData {
			// fmt.Printf("monkey %d. Items: %v\n", iMonkey, m.items)
			for _, item := range m.items {
				m.ItemCount++

				// update worry level
				// fmt.Println("item:", item)
				newWorryLevel := m.operation(*item)
				// fmt.Println(" new worry level:", item, newWorryLevel)

				// test for which monkey to throw item
				var toThrow int
				// mod := new(big.Int).Mod(newWorryLevel, big.NewInt(int64(m.test.divisor)))
				// if mod.Cmp(big.NewInt(0)) == 0 {
				if newWorryLevel.IsDivisibleByPrime(m.test.divisor) {
					// fmt.Println(newWorryLevel, m.test.divisor)
					//panic("err")
					toThrow = m.test.ifTrueMonkeyId
				} else {
					toThrow = m.test.ifFalseMonkeyId
				}

				// apply throw
				monkeysData[toThrow].items = append(monkeysData[toThrow].items, &newWorryLevel)
				// fmt.Println("oi", m.items, iItem, item, newWorryLevel)
			}

			// clear items that were throwed to other monkeys
			m.items = make([]*Item, 0)
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
