package slidingwindow

func countAnagramOccurences(str string, pattern string) int {

	var patFreq [26]byte
	k := len(pattern)
	// calculate pattern freq
	for i := range k {
		patFreq[pattern[i]-'a']++
	}

	var strFreq [26]byte
	// first window
	for i := range k {
		strFreq[str[i]-'a']++
	}

	result := 0

	if patFreq == strFreq {
		result++
	}

	n := len(str)
	// other window
	for i := 1; i < n-k+1; i++ {
		strFreq[str[i-1]-'a']--   // decrease freq
		strFreq[str[i+k-1]-'a']++ // increase freq

		if patFreq == strFreq {
			result++
		}
	}

	return result

}
