package hashmap

import (
	"fmt"
	"math"
	"slices"
)

func bucketSort(arr []float64) []float64 {

	n := len(arr)
	res := []float64{}

	if n == 0 {
		return res
	}

	maxElement := findMax(arr)

	var maxIndex int = int(maxElement * float64(n))

	bucketSize := maxIndex + 1

	// Create bucket list
	buckets := make([][]float64, bucketSize)
	// for i := range bucketSize {
	// 	buckets[i] = []float64{}
	// }

	// Insert elements into the bucket
	for _, val := range arr {
		var index int = int(val * float64(n))

		// buckets[index] = append(buckets[index], val)

		// Use Insertion sort
		targetIndex := 0
		for targetIndex < len(buckets[index]) && val > buckets[index][targetIndex] {
			targetIndex++
		}

		buckets[index] = slices.Insert(buckets[index], targetIndex, val)
	}

	// Sort the buckets
	// for i := range bucketSize {
	// 	sort.Float64s(buckets[i])
	// }

	// Store the elements into result list
	for i := range bucketSize {
		fmt.Printf("Bucket[%d] -> %v", i, buckets[i])
		fmt.Println()

		res = append(res, buckets[i]...)
	}

	return res
}

func findMax(arr []float64) float64 {
	max := -(math.MaxFloat64)

	for _, val := range arr {
		max = math.Max(max, val)
	}

	return max

}
