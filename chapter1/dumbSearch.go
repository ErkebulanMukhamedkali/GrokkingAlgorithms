package chapter1

func DumbSearch[T ~string | ~int](item T, list []T) (int, bool) {
	for i := range list {
		if item == list[i] {
			return i, true
		}
	}
	return 0, false
}
