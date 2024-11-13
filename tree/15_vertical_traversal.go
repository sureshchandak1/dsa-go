package tree

import (
	"reflect"
	"sort"
)

type VerticalPair struct {
	node               *TreeNode
	horizontalDistance int
	level              int
}

func verticalTraversal(root *TreeNode) []int {

	result := []int{}

	if root == nil {
		return result
	}

	nodes := make(map[int]map[int][]int, 0)

	queue := make([]*VerticalPair, 0)
	queue = append(queue, &VerticalPair{node: root, horizontalDistance: 0, level: 0})

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front

		node := front.node
		hd := front.horizontalDistance
		level := front.level

		// insert in map
		_, exists := nodes[hd]
		if !exists {
			nodes[hd] = make(map[int][]int)
			nodes[hd][level] = []int{}
		}
		nodes[hd][level] = append(nodes[hd][level], node.Data)

		if node.Left != nil {
			queue = append(queue, &VerticalPair{node: node.Left, horizontalDistance: hd - 1, level: level + 1})
		}

		if node.Right != nil {
			queue = append(queue, &VerticalPair{node: node.Right, horizontalDistance: hd + 1, level: level + 1})
		}
	}

	keys := reflect.ValueOf(nodes).MapKeys()
	keyList := []int{}
	for _, key := range keys {
		keyList = append(keyList, key.Interface().(int))
	}

	sort.Slice(keyList, func(i, j int) bool {
		return keyList[i] < keyList[j]
	})

	for _, hd := range keyList {
		for _, value := range nodes[hd] {
			result = append(result, value...)
		}
	}

	return result

}
