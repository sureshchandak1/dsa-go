package queue

import "container/list"

func firstNonRepeating(str string) string {

	count := make(map[byte]int)

	queue := list.New()

	result := ""

	for i := 0; i < len(str); i++ {
		ch := str[i]

		count[ch]++

		queue.PushBack(ch)

		for queue.Len() > 0 {

			if count[queue.Front().Value.(byte)] > 1 {
				queue.Remove(queue.Front())
			} else {
				result = result + string(queue.Front().Value.(byte))
				break
			}
		}

		if queue.Len() == 0 {
			result = result + "#"
		}
	}

	return result
}
