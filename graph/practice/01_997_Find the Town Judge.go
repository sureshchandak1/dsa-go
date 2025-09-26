package practice

func findJudgeOptimized(n int, trust [][]int) int {

	lenTrust := len(trust)

	if lenTrust < n-1 {
		return -1
	}

	degree := make([]int, n+1)

	for i := range lenTrust {
		degree[trust[i][0]]--
		degree[trust[i][1]]++
	}

	for i := 1; i <= n; i++ {
		if degree[i] == n-1 {
			return i
		}
	}

	return -1
}

func findJudge(n int, trust [][]int) int {

	lenTrust := len(trust)

	if lenTrust < n-1 {
		return -1
	}

	inDegree := make([]int, n+1)
	outDegree := make([]int, n+1)

	for i := 0; i < lenTrust; i++ {
		outDegree[trust[i][0]]++
		inDegree[trust[i][1]]++
	}

	for i := 1; i <= n; i++ {
		if outDegree[i] == 0 && inDegree[i] == n-1 {
			return i
		}
	}

	return -1
}
