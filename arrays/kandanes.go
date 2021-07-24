package main

func FindMaxSumOfContiguousSubarry(array []int) int {
	totalSum := 0
	maxSoFar := array[0]

	for _, num := range array {
		totalSum += num

		if totalSum > maxSoFar {
			maxSoFar = totalSum
		}

		if totalSum < 0 {
			totalSum = 0
		}
	}

	return maxSoFar
}
