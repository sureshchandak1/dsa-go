package queue

import "container/list"

func sumOfKsubArray(arr []int, windowSize int) int {

	maxi := list.New()
	mini := list.New()

	// Process first window
	for i := 0; i < windowSize; i++ {

		// Remove all previous smaller that are elements are useless.
		for maxi.Len() > 0 && arr[maxi.Back().Value.(int)] <= arr[i] {
			maxi.Remove(maxi.Back())
		}

		// Remove all previous greater elements that are useless.
		for mini.Len() > 0 && arr[mini.Back().Value.(int)] >= arr[i] {
			mini.Remove(mini.Back())
		}

		maxi.PushBack(i)
		mini.PushBack(i)
	}

	var ans = 0
	ans = ans + arr[maxi.Front().Value.(int)] + arr[mini.Front().Value.(int)]

	for i := windowSize; i < len(arr); i++ {

		// Remove all elements which are out of this window
		for maxi.Len() > 0 && i-maxi.Front().Value.(int) >= windowSize {
			maxi.Remove(maxi.Front())
		}
		for mini.Len() > 0 && i-mini.Front().Value.(int) >= windowSize {
			mini.Remove(mini.Front())
		}

		// addition new element
		// Remove all previous smaller that are elements are useless.
		for maxi.Len() > 0 && arr[maxi.Back().Value.(int)] <= arr[i] {
			maxi.Remove(maxi.Back())
		}

		// Remove all previous greater elements that are useless.
		for mini.Len() > 0 && arr[mini.Back().Value.(int)] >= arr[i] {
			mini.Remove(mini.Back())
		}

		maxi.PushBack(i)
		mini.PushBack(i)

		ans = ans + arr[maxi.Front().Value.(int)] + arr[mini.Front().Value.(int)]
	}

	return ans

}
