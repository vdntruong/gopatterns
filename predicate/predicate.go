package main

type Predicate[T any] func(T) bool

func Filter[T any](items []T, predicate Predicate[T]) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
