package gu

// Includes checks if a given value is present in a slice
func Includes[T comparable](slice []T, value T) bool {
	for _, e := range slice {
		if e == value {
			return true
		}
	}
	return false
}

// SameOrZero checks if every value in the slice is equal or the "zero" value
func SameOrZero[T comparable](args ...T) (bool, T) {
	zero := *new(T)
	vals := Filter(args, func(i T) bool {
		return i != zero
	})
	if len(vals) == 0 {
		return true, zero
	}
	// if all args are all the same, return the value
	isSame := true
	for i := 1; i < len(vals); i++ {
		if vals[i-1] != vals[i] {
			isSame = false
			break
		}
	}
	if isSame {
		return true, vals[0]
	}
	// if there are multiple values, return the zero value
	return false, zero
}

// Filter returns a new slice with only the values that pass the filter
func Filter[T comparable](slice []T, filter func(T) bool) []T {
	var result []T
	for _, e := range slice {
		if filter(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map iterates over a slice and returns a new slice with the results of the function
func Map[T any, R any](slice []T, mapper func(T) R) []R {
	var result []R
	for _, e := range slice {
		result = append(result, mapper(e))
	}
	return result
}

// Reduce iterates over a slice and returns a single value
func Reduce[T any, R any](slice []T, reducer func(R, T) R, initial R) R {
	for _, e := range slice {
		initial = reducer(initial, e)
	}
	return initial
}

// P return a pointer to the given value, useful when you need pointers to constants
func P[T any](v T) *T { return &v }

// D dereferences a pointer, if the pointer is nil, it will return a new value
func D[T any](v *T) T {
	if v == nil {
		v = new(T)
	}
	return *v
}
