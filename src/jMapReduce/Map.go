package jMapReduce

import (
	"github.com/jeltjongsma/jMapReduce/src/types"
)

type MapFunc[T, U any] func(T) U

func Map[T, U any](s JSlice[T], f MapFunc[T, U]) JSlice[U] {
	newJSlice := make([]U, s.size, s.capacity)
	for i, elem := range s {
		newJSlice[i] = f(elem)
	}
	return newJSlice
}

// func (s JSlice[T]) Map[U any](f MapFunc[T, U]) JSlice[U] {
// 	ret := NewJSlice[U](s.n_partitions, s.capacity)
// 	// Loop over all partitions
// 	for i, partition := range s.partitions {
// 		elements := make([]T, partition.size, partition.capacity)
// 		// Map each element
// 		for j, element := range partition.elements {
// 			elements[j] = f(element)
// 		}
// 		// Create new partition
// 		ret.partitions[i] = Partition[T]{
// 			elements: 	elements,
// 			size:		partition.size,
// 			capacity:	partition.capacity,
// 		}
// 	}
// 	return ret
// }