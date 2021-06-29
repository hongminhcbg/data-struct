package alv

import "fmt"

func leftRotate(root *Node)  {
	return
}

func clone(n *Node) *Node {
	if n == nil {
		return nil
	}

	return &Node{
		Value:  n.Value,
		Left:   n.Left,
		Right:  n.Right,
		Height: n.Height,
	}
}

func rightRotate(root *Node)  {
	fmt.Printf("root = %d, root.Left = %d, root.left.left = %d\n", root.Value, root.Left.Value, root.Left.Left.Value)

	oldRoot := clone(root)
	fmt.Printf("[db] Before ratate, root %v, root.Left = %v\n", oldRoot, root.Left)
	root = root.Left
	oldRoot.Left = root.Right
	root.Right = oldRoot
	fmt.Printf("[db] After ratate, root %v, root.Left = %v\n", root, root.Right)

	root.Height = max(getHeight(root.Left), getHeight(root.Right)) + 1
	oldRoot.Height = max(getHeight(oldRoot.Right), getHeight(oldRoot.Left)) + 1
	return
}
