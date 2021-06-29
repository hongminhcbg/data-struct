package heap

import "fmt"

var values []int
var length = 0

func Init(arr []int, N int) {
	for i := N/2; i > 0; i-- {
		maxHeap(arr, i, N)
	}

	values = make([]int, 0)
	values = append(values, arr...)
	length = N
}

func Pop() int {
	if length == 1 {
		panic("queue is have no data")
	}

	max := values[1]
	values[1] = values[length-1]
	length--
	values = values[:length]
	maxHeap(values, 1, length)
	return max
}
func SetVal(index int, val int)  {
	if values[index] > val {
		panic("not support")
	}

	values[index] = val
	for {
		if index == 1 {
			return
		}
		if values[index/2] > values[index] {
			return
		}

		temp := values[index/2]
		values[index/2] = values[index]
		values[index] = temp
		index = index / 2
	}
}

func Push(x int) {
	values = append(values, -1)
	length++
	SetVal(length-1, x)
}

func Print()  {
	fmt.Println(values)
}

func maxHeap(arr []int, i int, N int) {
	left := 2 * i
	right := 2 * i + 1
	largest := i
	if left < N && arr[largest] < arr[left] {
		largest = left
	}

	if right < N && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		temp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp
		maxHeap(arr, largest, N)
	}
}
