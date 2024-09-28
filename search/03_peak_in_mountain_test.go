package search

import (
	"fmt"
	"testing"
)

func TestPeakMountainIndex(t *testing.T) {

	arr := []int{1, 3, 5, 7, 10, 6, 4, 2}

	fmt.Println("Mountain Peak:", peakIndexInMountainArray(arr))

}
