package algotithms_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KadanesAlgorithm_ForAllPositiveNumbers(t *testing.T){
	sampleArray := []int{ 1, 4, 6, 2, 8}

	a := FindMaxSumOfContiguousSubarry(sampleArray)
	assert.Equal(t, 21, a)
}

func Test_KadanesAlgorithm_ForAllNegativeNumbers(t *testing.T){
	sampleArray := []int{ -11, -4, -6, -2, -8, -1}

	a := FindMaxSumOfContiguousSubarry(sampleArray)
	assert.Equal(t, -1, a)
}

func Test_KadanesAlgorithm_ForMixNumbers(t *testing.T){
	sampleArray := []int{ 5, 4, -10, 10, 1}

	a := FindMaxSumOfContiguousSubarry(sampleArray)
	assert.Equal(t, 11, a)
}