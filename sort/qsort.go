package sort

func QuickSort(array []int) []int {
	qsortRecursive(array, 0, len(array)-1)
	return array
}

func qsortRecursive(array []int, left int, right int) []int {
	if left > right {
		return array
	}
	i := left
	j := right
	pivot := array[(left+right)/2]
	for i <= j {
		for array[i] < pivot {
			i++
		}
		for array[j] > pivot {
			j--
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}
	qsortRecursive(array, left, j)
	qsortRecursive(array, i, right)
	return array
}
