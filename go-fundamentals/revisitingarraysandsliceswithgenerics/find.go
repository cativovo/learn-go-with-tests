package revisitingarraysandsliceswithgenerics

func Find[T comparable](s []T, cb func(v T) bool) (T, bool) {
	for _, v := range s {
		if cb(v) {
			return v, true
		}
	}

	var defaultValue T
	return defaultValue, false
}
