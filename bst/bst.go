package bst

import (
	"fmt"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
}

type BST interface {
	Insert(int)
	Find(int) bool
	Delete(int)
}

func NewNode(x int) Node {
	return Node{
		Value: x,
		Left:  nil,
		Right: nil,
	}
}

//
//func Init(root Node) BST {
//	return &root
//}

func insert(root *Node, x int)  {
	if x<= root.Value {
		if root.Left == nil {
			node := NewNode(x)
			root.Left = &node
			return
		}

		insert(root.Left, x)
		return
	}

	if root.Right == nil {
		node := NewNode(x)
		root.Right = &node
		return
	}

	insert(root.Right,x)
}

func (t *Node) Insert(x int) {
	insert(t, x)
}

func print(root *Node)  {
	if root.Left != nil {
		print(root.Left)
	}

	fmt.Print(root.Value, " ")

	if root.Right != nil {
		print(root.Right)
	}
}

func (t *Node) Print() {
	printx("", t, true)
}

func find(root *Node, x int) bool {
	if root == nil {
		return false
	}

	if x == root.Value {
		return true
	}

	if x < root.Value {
		return find(root.Left, x)
	}

	return find(root.Right, x)
}

func (t *Node) Find(x int) bool {
	return find(t, x)
}

func findNode(root *Node, x int) *Node {
	if root == nil {
		return nil
	}

	if root.Value == x {
		return root
	}

	if x < root.Value {
		return findNode(root.Left, x)
	}

	return findNode(root.Right, x)
}

func findLeftMostOfNode(root *Node) *Node {
	if root.Left == nil {
		return root
	}

	return findLeftMostOfNode(root.Left)
}

func findRightMostOfNode(root *Node) *Node {
	if root.Right == nil {
		return root
	}

	return findRightMostOfNode(root.Right)
}

func deleteNodeWithSingleChild(myParent *Node, me *Node, target int) {
	if me.Value == target {
		var onlyMyChild *Node = nil
		if me.Left != nil {
			onlyMyChild = me.Left
		} else {
			onlyMyChild = me.Right
		}

		if myParent.Left != nil && myParent.Left.Value == me.Value {
			myParent.Left = onlyMyChild
		} else {
			myParent.Right = onlyMyChild
		}

		me = nil
		return
	}

	if target < me.Value {
		deleteNodeWithSingleChild(me, me.Left,target)
		return
	}

	deleteNodeWithSingleChild(me, me.Right, target)
	return
}

func deleteRoot(root *Node)  {
	if root.Left != nil {
		rightMostOfLeft := findRightMostOfNode(root.Left)
		root.Value = rightMostOfLeft.Value
		deleteNodeWithSingleChild(root, root.Left, root.Value)
		return
	}

	if root.Right != nil {
		leftMostOfRight := findLeftMostOfNode(root.Right)
		root.Value = leftMostOfRight.Value
		deleteNodeWithSingleChild(root, root.Right, root.Value)
		return
	}

	panic("after delete bst is nil")
}


func (t *Node) Delete(x int)  {
	if x == t.Value {
		deleteRoot(t)
		return
	}

	nodeNeedDelete :=  findNode(t, x)
	if nodeNeedDelete == nil {
		return
	}

	//case 1 node have no child
	if nodeNeedDelete.Right == nil && nodeNeedDelete.Left == nil {
		nodeNeedDelete = nil
		return
	}

	//case 2 node have 1 child
	if nodeNeedDelete.Right == nil || nodeNeedDelete.Left == nil  {
		deleteNodeWithSingleChild(nil, t, x)
		return
	}

	// case3 node have 2 child
	leftMostOfNodeRight := findLeftMostOfNode(nodeNeedDelete.Right)
	nodeNeedDelete.Value = leftMostOfNodeRight.Value
	deleteNodeWithSingleChild(nodeNeedDelete, nodeNeedDelete.Right, leftMostOfNodeRight.Value)
}

func printx(prefix string, root *Node, isLeft bool)  {
	if root == nil {
		return
	}

	if isLeft {
		fmt.Printf("%s\\--> %d\n", prefix,  root.Value)
	} else {
		fmt.Printf("%s---> %d\n", prefix, root.Value)
	}

	if isLeft {
		prefix = prefix + "    "
	} else {
		prefix = prefix + "|    "
	}

	printx(prefix, root.Right, false)
	printx(prefix, root.Left, true)
}
