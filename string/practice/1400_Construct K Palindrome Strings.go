package practice

func canConstruct(s string, k int) bool {

	n := len(s)

	if n < k {
		return false
	}

	if n == k {
		return true
	}

	count := make([]int, 26)
	for i := 0; i < n; i++ {
		count[s[i]-'a']++
	}

	oddCount := 0
	for i := 0; i < 26; i++ {
		if count[i]%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= k
}
