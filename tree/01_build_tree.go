package tree

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func buildTree() *TreeNode {

	fmt.Println("Enter the data: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}

	nodeValue, _ := strconv.Atoi(input)
	root := TreeNode{Data: nodeValue}

	if nodeValue == -1 {
		return nil
	}

	fmt.Printf("Enter data for inserting in left of %d:\n", nodeValue)
	root.Left = buildTree()

	fmt.Printf("Enter data for inserting in right of %d:\n", nodeValue)
	root.Right = buildTree()

	return &root
}

func buildTreeString(str string) *TreeNode {

	if len(str) == 0 || str[0] == 'N' {
		return nil
	}

	var ip []string = strings.Split(str, " ")
	size := len(ip)

	rootVal, _ := strconv.Atoi(ip[0])
	root := &TreeNode{Data: rootVal}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	i := 1
	for len(queue) > 0 && i < size {

		// get front from queue
		currNode := queue[0]
		queue = queue[1:] // remove front from queue

		currVal := ip[i]

		if currVal != "N" {
			leftVal, _ := strconv.Atoi(currVal)
			currNode.Left = &TreeNode{Data: leftVal}
			queue = append(queue, currNode.Left)
		}

		i++

		if i >= size {
			break
		}

		currVal = ip[i]

		if currVal != "N" {
			rightVal, _ := strconv.Atoi(currVal)
			currNode.Right = &TreeNode{Data: rightVal}
			queue = append(queue, currNode.Right)
		}

		i++
	}

	return root
}
