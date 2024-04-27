package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	var s Stack[int]
	t.Run("create", func(t *testing.T) {
		s = NewStack[int]()
		assert.Equal(t, s.Length(), 0)
	})

	t.Run("push", func(t *testing.T) {
        s.Push(1)

		assert.Equal(t, s.Peek(), 1)

        s.Push(2)

		assert.Equal(t, s.Peek(), 2)
	})

	t.Run("pop", func(t *testing.T) {
		fmt.Println(s)
		val := s.Pop()
		fmt.Println(s)
		cur := s.Peek()

		assert.Equal(t, val, 2)
		assert.Equal(t, cur, 1)
	})

}
