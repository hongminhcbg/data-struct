package heap

import (
	"fmt"
	"testing"
)

func Test_InitHeap(t *testing.T) {
	h := Newv2(3, 4, 5, 6, 7, 1, 3, 4, 5, 9, 3, 2)
	if h.Pop() != 9 {
		t.Error("error")
	}
	if h.Pop() != 7 {
		t.Error("error")
	}

	for h.Size() != 0 {
		fmt.Println(h.Pop())
	}

	for i := 1; i < 10; i++ {
		h.Insert(i)
	}

	for h.Size() != 0 {
		fmt.Println(h.Pop())
	}
}
