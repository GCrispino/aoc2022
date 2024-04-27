package types

type Pair[T any] struct {
	First, Second T
}

func NewPair[T any](first, second T) Pair[T]{
    return Pair[T]{
        First: first,
        Second: second,
    }
}

type Stack[T any] []T

func (s *Stack[T]) Pop() T {
    length := (*s).Length()

    val := (*s)[length - 1]

    *s = (*s)[:length - 1]

    return val
}

func (s *Stack[T]) Push(val T){
    *s = append(*s, val)
}

func (s Stack[T]) Peek() T {
    return s[s.Length() - 1]
}

func (s Stack[T]) Length() int {
    return len(s)
}

func NewStack[T any]() Stack[T] {
    return make(Stack[T], 0)
}
