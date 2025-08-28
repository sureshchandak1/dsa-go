package practice

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {

	ver1Arr := strings.Split(version1, ".")
	ver2Arr := strings.Split(version2, ".")

	n1 := len(ver1Arr)
	n2 := len(ver2Arr)

	p1, p2 := 0, 0
	for p1 < n1 || p2 < n2 {
		num1 := 0
		if p1 < n1 {
			num1, _ = strconv.Atoi(ver1Arr[p1])
		}
		num2 := 0
		if p2 < n2 {
			num2, _ = strconv.Atoi(ver2Arr[p2])
		}

		if num1 != num2 {
			if num1 > num2 {
				return 1
			}
			return -1
		}

		p1++
		p2++
	}

	return 0
}

func compareVersionOptimized(version1 string, version2 string) int {

	n1 := len(version1)
	n2 := len(version2)

	p1, p2 := 0, 0
	for p1 < n1 || p2 < n2 {

		num1, num2 := 0, 0

		for p1 < n1 && version1[p1] != '.' {
			num1 = num1*10 + int(version1[p1]-'0')
			p1++
		}

		for p2 < n2 && version2[p2] != '.' {
			num2 = num2*10 + int(version2[p2]-'0')
			p2++
		}

		if num1 != num2 {
			if num1 > num2 {
				return 1
			}
			return -1
		}

		p1++ // skip the dot
		p2++ // skip the dot
	}

	return 0
}
