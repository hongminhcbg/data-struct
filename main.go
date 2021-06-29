package main

import (
	"data-struct/alv"
	"data-struct/bst"
	"fmt"
)

func demoBst() {
	fmt.Println("DEMO BST")
	root := bst.NewNode(50)

	bst := &root
	bst.Insert(25)
	bst.Insert(100)
	bst.Insert(10)
	bst.Insert(7)
	bst.Insert(69)
	bst.Insert(80)
	bst.Insert(12)
	bst.Insert(76)
	bst.Insert(14)
	bst.Insert(84)
	bst.Insert(1)

	bst.Print()
	//if bst.Find(45) {
	//	fmt.Println("45 Found")
	//} else {
	//	fmt.Println("45 Not found")
	//}
	//
	//if bst.Find(8) {
	//	fmt.Println("8 Found")
	//} else {
	//	fmt.Println("8 Not found")
	//}

	bst.Delete(80)
	bst.Print()

	bst.Delete(50)
	bst.Print()

}

func demoAVL() {
	node := alv.NewNode(50)
	node.Insert(25)
	node.Insert(10)
	//node.Insert(75)
	//node.Insert(85)
	//node.Insert(12)
	//node.Insert(19)
	//node.Insert(6)
	//node.Insert(3)
	//node.Insert(64)
	//node.Insert(86)
	//node.Insert(15)

	node.Print()
}
func main() {
	demoAVL()
}
