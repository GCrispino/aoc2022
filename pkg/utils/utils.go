package utils

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"

	big "github.com/ncw/gmp"
)

func ReadLines(filePath string) []string {
	lines := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func SumInts(vals []int) (sum int) {
	for _, val := range vals {
		sum += val
	}

	return
}

func BuildSetMapFromContainer[T comparable](container []T) map[T]struct{} {
	setMap := make(map[T]struct{})

	for _, item := range container {
		setMap[item] = struct{}{}
	}

	return setMap
}

func Map[T, U any](sl []T, fn func(T) U) []U {
	newSl := make([]U, len(sl))

	for i, val := range sl {
		newSl[i] = fn(val)
	}

	return newSl
}

func MapWithError[T, U any](sl []T, fn func(T) (U, error)) ([]U, error) {
	newSl := make([]U, len(sl))

	var err error
	for i, val := range sl {
		if newSl[i], err = fn(val); err != nil {
			return nil, err
		}
	}

	return newSl, nil
}

func CheckAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func SplitAndGetNthItemInt(s, sep string) int {
	spl := strings.Split(s, sep)
	x, err := strconv.Atoi(spl[1])
	CheckAndPanic(err)

	return x
}

var (
	bigIntZero  = big.NewInt(0)
	bigIntTwo   = big.NewInt(2)
	bigIntThree = big.NewInt(3)
)

// Get all prime factors of a given number n
func PrimeFactorsOrig(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// Get all prime factors of a given number n
func PrimeFactors(n *big.Int) (pfs []*big.Int) {
	blah := new(big.Int)
	// Get the number of 2s that divide n
	for blah.Mod(n, bigIntTwo).Int64() == 0 {
		pfs = append(pfs, bigIntTwo)
		n.Div(n, bigIntTwo)
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)

	i := big.NewInt(3)
	blah = new(big.Int)
	for blah.Mul(i, i).Cmp(n) <= 0 {
		// while i divides n, append i and divide n
		for blah.Mod(n, i).Cmp(bigIntZero) == 0 {
			pfs = append(pfs, new(big.Int).Set(i))
			n.Div(n, i)
		}
		i.Add(i, bigIntTwo)
	}

	// This condition is to handle the case when n is a prime number
	if n.Cmp(bigIntTwo) == 1 {
		pfs = append(pfs, n)
	}

	return
}
