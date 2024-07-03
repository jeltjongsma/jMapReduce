package	jMapReduce

type FilterFunc[T] func(T) bool

func Filter[T any](s JSlice[T], f FilterFunc[T]) JSlice[T] {
	newJSlice := make([]T, 0, s.capacity)
	for _, elem := range s {
		if f(elem) {
			newJSlice = append(newJSlice, elem)
		}
	}
	return newJSlice
} 