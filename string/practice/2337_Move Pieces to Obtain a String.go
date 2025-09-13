package practice

func movePiecesToTarget(start string, target string) bool {

	n := len(start)
	m := len(target)

	i, j := 0, 0

	for i < n || j < m {

		// skip all blanks in start
		for i < n && start[i] == '_' {
			i++
		}

		// skip all blanks in target
		for j < m && target[j] == '_' {
			j++
		}

		// count is same only if the string end at the same time
		if i == n || j == m {
			return i == n && j == m
		}

		// check false case
		if start[i] != target[j] {
			return false
		}
		if start[i] == 'L' && j > i {
			return false
		}
		if start[i] == 'R' && j < i {
			return false
		}

		i++
		j++
	}

	return true
}
