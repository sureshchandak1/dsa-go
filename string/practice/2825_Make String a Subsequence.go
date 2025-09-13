package practice

func canMakeSubsequence(str1 string, str2 string) bool {

	n := len(str1)
	m := len(str2)

	if m > n {
		return false
	}

	i, j := 0, 0

	for j < m && i < n {
		if str1[i] == str2[j] || str1[i] == str2[j]-1 || (str1[i] == 'z' && str2[j] == 'a') {
			j++
		}
		i++
	}

	return j == m
}
