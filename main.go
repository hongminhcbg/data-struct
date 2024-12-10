package main

import (
	"data-struct/alv"
	"data-struct/bst"
	"data-struct/heap"
	"data-struct/mapv2"
	"fmt"
	"time"
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
	// node.Insert(75)
	// node.Insert(85)
	// node.Insert(12)
	// node.Insert(19)
	// node.Insert(6)
	// node.Insert(3)
	// node.Insert(64)
	// node.Insert(86)
	// node.Insert(15)

	node.Print()
}

func demoHeap() {
	arr := []int{0, 1, 4, 3, 7, 8, 9, 10}
	heap.Init(arr, len(arr))
	fmt.Println("Heap init success")
	heap.Print()

	fmt.Println("Pop 3 elements")
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())

	fmt.Println("Heap after pop")
	heap.Print()

	fmt.Println("Replace 4 -> 15")
	heap.SetVal(2, 15)

	fmt.Println("Heap after replace")
	heap.Print()

	fmt.Println("Push 3 elements 8 90 2 to heap")
	heap.Push(8)
	heap.Push(90)
	heap.Push(2)

	fmt.Println("Heap after push")
	heap.Print()
}

func demomapv2() {
	const MAX = 100
	m := mapv2.NewMap()
	newMap := make(map[int]int)
	for i := 0; i < MAX; i++ {
		newMap[i] = i % 7
	}

	m.Load(newMap)
	fmt.Println(m.Get(3))

	// make 2000 goroutines to read map
	for i := 0; i < 2000; i++ {
		go func() {
			time.Sleep(time.Duration(i%MAX)*time.Millisecond + time.Duration(i/200)*time.Second)
			fmt.Println("go routine: ", i, m.Get(i%MAX))
		}()
	}

	// make internal 1s job set new data to map
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			newMap := make(map[int]int)
			for i := 0; i < MAX; i++ {
				newMap[i] = i % 63
			}

			m.Load(newMap)
			fmt.Println("Load new map success")
		}
	}()

	// sleep 10s to see the result
	time.Sleep(10 * time.Second)
}

func main() {
	// dfs.DFS(3, 10)
	demomapv2()
}
