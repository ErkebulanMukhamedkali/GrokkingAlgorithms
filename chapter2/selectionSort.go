package chapter2

func SelectionSort[T ~string | ~int](arr []T) []T {
	var newArr []T = make([]T, len(arr))
	for i := range arr {
		var smallest int = findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func findSmallest[T ~string | ~int](arr []T) int {
	var smallest T = arr[0]
	var smallestIndex int = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func SelectionSortV2[T ~string | ~int](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
