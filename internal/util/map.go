package util

func MapKeys[T, V comparable](m map[T]V) []T {
	var keys []T
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[T, V comparable](m map[T]V) []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
