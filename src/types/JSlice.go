package types

type JSlice[T any] []T

// Code for v0.2
// type JSlice[T any] struct {
// 	partitions 		[]Partition[T]
// 	n_partitions	int
// 	capacity		int 
// }

// type Partition[T any] struct {
// 	elements		[]T 
// 	size 			int 
// 	capacity 		int	
// 	id				int
// }

// // Initialize a JSlice 
// func NewJSlice[T any](n_partitions, capacity int) JSlice[T] {
// 	jSlice := JSlice[T]{
// 		partitions: 	make([]Partition[T], n_partitions, capacity),
// 		n_partitions:	n_partitions,
// 		capacity:   	capacity,
// 	}

// 	for i := range jSlice.partitions {
// 		jSlice.partitions[i] = Partition[T]{
// 			elements: make([]T, 0),
// 			size:     0,
// 			capacity: 0,
// 			id:       i,
// 		}
// 	}

// 	return jSlice
// }