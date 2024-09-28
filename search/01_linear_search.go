package search

func linearSearch(arr []int, item int) (bool, int, int) {

	for index, value := range arr {
		if value == item {
			return true, index, value
		}
	}

	return false, -1, -1
}
