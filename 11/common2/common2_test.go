package common2

import (
	"testing"

	big "github.com/ncw/gmp"
	"github.com/stretchr/testify/assert"

	"aoc-2022/pkg/utils"
)

func TestItem(t *testing.T) {
	var item *Item
	t.Run("new item from non-prime number", func(t *testing.T) {
		item = NewItem(big.NewInt(112359))
		expected := utils.Map([]int64{3, 13, 43, 67}, big.NewInt)
		assert.Equal(t, expected, item.factors)
	})

	t.Run("multiply with prime numbers", func(t *testing.T) {
		item.Mul(big.NewInt(17))
		expected := utils.Map([]int64{3, 13, 17, 43, 67}, big.NewInt)
		assert.Equal(t, expected, item.factors)
		e := item.Eval()
		assert.Equal(t, 0, e.Cmp(big.NewInt(1910103)))

		item.Mul(big.NewInt(2))
		expected = utils.Map([]int64{2, 3, 13, 17, 43, 67}, big.NewInt)
		assert.Equal(t, expected, item.factors)
		e = item.Eval()
		assert.Equal(t, 0, e.Cmp(big.NewInt(3820206)))

		item.Mul(big.NewInt(71))
		expected = utils.Map([]int64{2, 3, 13, 17, 43, 67, 71}, big.NewInt)
		assert.Equal(t, expected, item.factors)
		e = item.Eval()
		assert.Equal(t, 0, e.Cmp(big.NewInt(271234626)))

		item.Mul(big.NewInt(13))
		expected = utils.Map([]int64{2, 3, 13, 13, 17, 43, 67, 71}, big.NewInt)
		assert.Equal(t, expected, item.factors)
		e = item.Eval()
		assert.Equal(t, 0, e.Cmp(big.NewInt(3526050138)))
	})

	t.Run("checks if it is divisible by prime", func(t *testing.T) {
		assert.Equal(t, false, item.IsDivisibleByPrime(19))

		assert.Equal(t, true, item.IsDivisibleByPrime(13))

		assert.Equal(t, true, item.IsDivisibleByPrime(17))

		assert.Equal(t, false, item.IsDivisibleByPrime(73))
	})

	t.Run("add numbers", func(t *testing.T) {
		item.Add(big.NewInt(10))
		assert.Equal(t, 0, big.NewInt(3526050148).Cmp(item.Eval()))

		item.Add(big.NewInt(7))
		assert.Equal(t, 0, big.NewInt(3526050155).Cmp(item.Eval()))
	})
}
