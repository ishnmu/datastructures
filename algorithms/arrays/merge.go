package arrays

func Merge(array1 []int, array2 []int) []int {
	mergedSortedList := make([]int, (len(array1) + len(array2)))
	indexPointer1 := 0
	indexPointer2 := 0
	indexPointerMerged := 0

	for len(array1) > indexPointer1 && len(array2) > indexPointer2 {
		if array1[indexPointer1] < array2[indexPointer2] {
			mergedSortedList[indexPointerMerged] = array1[indexPointer1]
			indexPointer1++
		} else {
			mergedSortedList[indexPointerMerged] = array2[indexPointer2]
			indexPointer2++
		}
		indexPointerMerged++
	}

	for indexPointer1 < len(array1) {
		mergedSortedList[indexPointerMerged] = array1[indexPointer1]
		indexPointer1++
		indexPointerMerged++
	}

	for indexPointer2 < len(array2) {
		mergedSortedList[indexPointerMerged] = array2[indexPointer2]
		indexPointer2++
		indexPointerMerged++
	}

	return mergedSortedList
}
