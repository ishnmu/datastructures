package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge_ForTwoSorted_EqualSizedArray(t *testing.T) {
	array1 := []int{1, 12, 23, 44, 85}
	array2 := []int{12, 20, 25, 40, 56}

	mergedSortedArray := Merge(array1, array2)
	assert.Equal(t, []int{1, 12, 12, 20, 23, 25, 40, 44, 56, 85}, mergedSortedArray)
}

func TestMerge_ForTwoSorted_UnEqualSizedArray(t *testing.T) {
	array1 := []int{1, 12, 23, 44, 85}
	array2 := []int{12, 20, 25, 40, 56, 100, 1000, 10001}

	mergedSortedArray := Merge(array1, array2)
	assert.Equal(t, []int{1, 12, 12, 20, 23, 25, 40, 44, 56, 85, 100, 1000, 10001}, mergedSortedArray)
}
