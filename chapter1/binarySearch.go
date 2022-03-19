package chapter1

func BinarySearch[T ~string | ~int](item T, list []T) (int, bool) {
	var low, high int = 0, len(list) - 1
	for low <= high {
		var mid int = (low + high) / 2
		var guess T = list[mid]
		switch {
		case guess == item:
			return mid, true
		case guess > item:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return 0, false
}
