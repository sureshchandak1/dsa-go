package practice

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {

	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	freqMap := make(map[int]int)
	for _, element := range hand {
		freqMap[element]++
	}

	sort.Ints(hand)

	for _, startCard := range hand {

		if freqMap[startCard] == 0 {
			continue
		}

		for freqMap[startCard-1] > 0 {
			startCard--
		}

		for i := range groupSize {

			card := startCard + i

			if freqMap[card] == 0 {
				return false
			}

			freqMap[card]--

			if freqMap[card] == 0 {
				delete(freqMap, card)
			}
		}
	}

	return true
}
