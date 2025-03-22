package graph

import "container/list"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {

	made := []string{}

	// Create graph
	adj := make(map[string][]string)
	inDegree := make(map[string]int)

	for index, recipe := range recipes {
		for _, ing := range ingredients[index] {
			adj[ing] = append(adj[ing], recipe)
			inDegree[recipe]++
		}
	}

	queue := list.New()
	for _, supply := range supplies {
		queue.PushBack(supply)
	}

	for queue.Len() != 0 {
		front := queue.Front()
		queue.Remove(front)

		for _, next := range adj[front.Value.(string)] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue.PushBack(next)
				made = append(made, next)
			}
		}
	}

	return made
}
