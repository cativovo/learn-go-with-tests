package revisitingarraysandsliceswithgenerics

func Reduce[T any, V any](slice []T, cb func(accumulator V, currentValue T) V, initialValue V) V {
	result := initialValue
	for _, v := range slice {
		result = cb(result, v)
	}

	return result
}
