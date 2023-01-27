package util

func SliceContains[T comparable](xs []T, y T) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}
