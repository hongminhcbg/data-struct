package heap

import "fmt"

type IHeap interface {
	Size() int
	Insert(x int)
	Pop() int
	Top() int
}

type heapImpl struct {
	size int
	arr  []int
}

// maxHeapify check arr[i] greater node left and right
// if true do nothing
// else swap i with left or right, after that call recursive go down
// this func must be call N/2 times when init heap
func (h *heapImpl) maxHeapify(i int) {
	left := 2 * i
	right := 2*i + 1
	greatest := i
	if left <= h.size && h.arr[left] > h.arr[greatest] {
		greatest = left
	}

	if right <= h.size && h.arr[right] > h.arr[greatest] {
		greatest = right
	}

	if greatest != i {
		h.arr[greatest], h.arr[i] = h.arr[i], h.arr[greatest]
		h.maxHeapify(greatest)
	}
}

func Newv2(arr ...int) IHeap {
	size := len(arr)
	marr := make([]int, size+1)

	for i, a := range arr {
		marr[i+1] = a
	}

	h := &heapImpl{
		size: size,
		arr:  marr,
	}
	for i := h.size / 2; i > 0; i-- {
		h.maxHeapify(i)
	}

	fmt.Println(h.arr)
	return h
}

func (h *heapImpl) Size() int {
	return h.size
}

func (h *heapImpl) Top() int {
	if h.size == 0 {
		panic("heap is empty, please check size before call top() or pop()")
	}

	return h.arr[1]
}

// Pop this func return max value and remove it
// solution is move a small element from end off queue and call func heapify
func (h *heapImpl) Pop() int {
	if h.size == 0 {
		panic("heap is empty, please check size before call top() or pop()")
	}

	max := h.arr[1]
	if h.size == 1 {
		// queue have only one element, can't get small element
		h.size = 0
		h.arr = []int{}
		return max
	}

	// move a small element to top
	h.arr[1] = h.arr[h.size]

	// remove last element
	h.arr = h.arr[:h.size]
	h.size -= 1
	h.maxHeapify(1)
	return max
}

// Insert x to heap, solution is, we will add x to and of arr
// compare x with his parent, if x > its parent swap and to up
// only its parent, don't care left or same level node
func (h *heapImpl) Insert(x int) {
	if h.size == 0 {
		h.size = 1
		h.arr = []int{0, x}
		return
	}

	h.arr = append(h.arr, x)
	h.size += 1
	now := h.size
	for {
		parent := now / 2
		if parent == 0 {
			break
		}

		if h.arr[now] <= h.arr[parent] {
			break
		}

		h.arr[now], h.arr[parent] = h.arr[parent], h.arr[now]
		now = parent
	}
}
