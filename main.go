package main

import (
	"fmt"

	"data-struct/bst"
)

func main() {
	fmt.Println("Hello world")
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
