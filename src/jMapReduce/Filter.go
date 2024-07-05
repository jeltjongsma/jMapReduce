package	jMapReduce

import (
	"github.com/jeltjongsma/jMapReduce/src/types"
)

type FilterFunc[T any] func(T) bool

func Filter[T any](s types.JSlice[T], f FilterFunc[T]) types.JSlice[T] {
	newJSlice := make([]T, 0, s.Capacity())
	for _, elem := range s {
		if f(elem) {
			newJSlice = append(newJSlice, elem)
		}
	}
	return newJSlice
} 