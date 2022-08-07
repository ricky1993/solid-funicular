package array

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Stack[T]) Top() *T {
	if len(*s) > 0 {
		return &(*s)[len(*s)-1]
	}
	return nil
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Print() {
	for _, elem := range *s {
		fmt.Print(elem)
		fmt.Print(" ")
	}
	fmt.Println("")
}
