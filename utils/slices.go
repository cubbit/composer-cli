package utils

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Find[T any](slice []T, findFn func(cmp T) bool) (T, bool) {
	for _, el := range slice {
		if findFn(el) {
			return el, true
		}
	}
	var result T
	return result, false
}

func Count[T any](slice []T, findFn func(cmp T) bool) int {
	count := 0
	for _, el := range slice {
		if findFn(el) {
			count++
		}
	}
	return count
}
