package sort

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	arrayLeft := MergeSort(array[:len(array)/2])
	arrayRight := MergeSort(array[len(array)/2:])
	return merge(arrayLeft, arrayRight)
}

func merge(a1 []int, a2 []int) (result []int) {
	for len(a1) > 0 && len(a2) > 0 {
		if a1[0] < a2[0] {
			result = append(result, a1[0])
			a1 = a1[1:]
		} else {
			result = append(result, a2[0])
			a2 = a2[1:]
		}
	}

	if len(a1) > 0 {
		result = append(result, a1...)
	}

	if len(a2) > 0 {
		result = append(result, a2...)
	}

	return result
}
