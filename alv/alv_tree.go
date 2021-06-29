package alv

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
	Height int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func NewNode(x int) *Node {
	return &Node{
		Value: x,
		Right: nil,
		Left: nil,
		Height: 1,
	}
}

type AVLTree interface {
	Search(x int) bool
	Insert(x int)
	Max() int
	Min() int
	Delete(x int)
}

func print(prefix string, root *Node, isLeft bool)  {
	if root == nil {
		return
	}

	if isLeft {
		fmt.Printf("%s\\--> %d(%d)\n", prefix,  root.Value, root.Height)
	} else {
		fmt.Printf("%s---> %d(%d)\n", prefix, root.Value, root.Height)
	}

	if isLeft {
		prefix = prefix + "    "
	} else {
		prefix = prefix + "|    "
	}

	print(prefix, root.Right, false)
	print(prefix, root.Left, true)
}

func (t *Node) Print() {
	print("", t, true)
}

func getHeight(n *Node) int {
	if n == nil {
		return 0
	}

	return n.Height
}

func balance(root *Node, x int)  {
	diff := getHeight(root.Left) - getHeight(root.Right)

	if diff > 1 {
		// left higher than right
		fmt.Printf("node[%d] need to balance\n", root.Value)
		if x < root.Left.Value {
			// Left Left case
			rightRotate(root)
			return
		}

		//left right case
		leftRotate(root.Left)
		rightRotate(root)
		return
	}

	if diff < -1 {
		// right node higher than left node
		if x > root.Right.Value {
			// case right right
			leftRotate(root)
			return
		}

		// right left case
		rightRotate(root.Right)
		leftRotate(root)
		return
	}
}

func insert(root *Node, x int) {
		defer balance(root, x)
		if x <= root.Value {
			if root.Left == nil {
				root.Left = NewNode(x)
				root.Height = max(2, getHeight(root.Right))
				return
			}

			insert(root.Left, x)
			root.Height = max(getHeight(root.Left), getHeight(root.Right)) + 1
			return
		}

		if root.Right == nil {
			root.Right = NewNode(x)
			root.Height = max(2, getHeight(root.Left))
			return
		}

		insert(root.Right, x)
		root.Height = max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func (t *Node) Insert(x int) {
	insert(t, x)
}


