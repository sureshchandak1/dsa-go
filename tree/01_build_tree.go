package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	buildTree()
}

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
	root := TreeNode{data: nodeValue}

	if nodeValue == -1 {
		return nil
	}

	fmt.Printf("Enter data for inserting in left of %d:\n", nodeValue)
	root.left = buildTree()

	fmt.Printf("Enter data for inserting in right of %d:\n", nodeValue)
	root.right = buildTree()

	return &root
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}
