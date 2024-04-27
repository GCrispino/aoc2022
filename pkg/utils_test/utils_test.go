package utilstest

import (
	"context"
	"fmt"
	"testing"
	"time"

	big "github.com/ncw/gmp"
	"github.com/stretchr/testify/assert"

	"aoc-2022/pkg/utils"
)

func TestPrimeFactors(t *testing.T) {
	testPrime := func(n *big.Int, expectedFactors ...*big.Int) func(t *testing.T) {
		return func(t *testing.T) {

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			ch := make(chan struct{})
			go func() {
				factors := utils.PrimeFactors(n)
				assert.Equal(t, expectedFactors, factors)

				ch <- struct{}{}
			}()

			select {
			case <-ctx.Done():
				t.Fatal(ctx.Err())
			case <-ch:
			}
		}
	}
	t.Run("a prime number has a single factor", testPrime(big.NewInt(2), big.NewInt(2)))

	t.Run("another prime number has a single factor", testPrime(big.NewInt(19), big.NewInt(19)))

	t.Run(
		"non-prime number has multiple factors",
		testPrime(big.NewInt(777), (utils.Map([]int64{3, 7, 37}, big.NewInt))...),
	)

	t.Run(
		"other non-prime number has multiple factors",
		testPrime(big.NewInt(112359), (utils.Map([]int64{3, 13, 43, 67}, big.NewInt))...),
	)
}

func TestPrimeFactorsOrig(t *testing.T) {
	t.Run("a prime number has a single factor", func(t *testing.T) {

		n := 2
		fmt.Println(n)
		factors := utils.PrimeFactorsOrig(n)
		expectedFactors := []int{2}
		assert.Equal(t, expectedFactors, factors)
	})

	t.Run("non-prime number has multiple factors", func(t *testing.T) {

		n := 777
		fmt.Println(n)
		factors := utils.PrimeFactorsOrig(n)
		expectedFactors := []int{3, 7, 37}
		assert.Equal(t, expectedFactors, factors)
	})
}
