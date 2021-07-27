package algotithms_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	array := []int{1, 12, 9, 23, 2, 3 ,44, 8, 85}

	sorted := MergeSort(array)
	assert.Equal(t, []int{1,2,3, 8, 9, 12,23,44, 85}, sorted)
}
