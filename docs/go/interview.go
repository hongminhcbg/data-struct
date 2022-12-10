package main

import (
	"fmt"
	"time"
	"unsafe"

	"sync"
)

func forRange() {
	a := []int{2, 4, 5}
	for _, x := range a {
		go func() {
			fmt.Println(x)
		}()
	}

	time.Sleep(1 * time.Second)
	// outout is 5, 5, 5
}

func deadlock() {
	c := make(chan int)
	c <- 3

	// do something
	x := <-c
	fmt.Println(x)
}

type richString string

func xxx(y richString) {
	fmt.Println(y)
}

func untypeConst() {
	// s := "abc"
	// xxx(s) // => error because s type string
}

func deferStack() {
	defer func() {
		fmt.Println("123")
	}()
	defer func() {
		fmt.Println("456")
	}()

	fmt.Println("789")
	// output is 789456123
}

func rune() {
	x := 'c'
	fmt.Printf("%T %d\n", x, unsafe.Sizeof(x))

	// default type is rune, size 4 byte, that is unicode32
}

func slice() {
	a := []int{3, 4, 5, 6, 7}
	b := a[1:3]
	a[2] = 9
	fmt.Println(b)
	// will print 4, 9 because b only reference to a
}

func slicev2() {
	a := make([]int, 0, 2)
	a = append(a, 1, 2)
	b := a[:]

	a = append(a, 3)
	a[0] = 5
	fmt.Println(b)
	// will print 1, 2 because b only reference to a
	// and when append 3 to 3, a will be resliced
}

func raceCondition() {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			x = x + 1
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(x)
}

func main() {
	fmt.Println("hello world")
	// forRange()
	// deadlock()
	//deferStack()
	//rune()
	// slice()
	// slicev2()
	//raceCondition()

	//memLeak()
	updateMapInsideLoop()
}

func xxx2() {
	x := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			x = x + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(x)
}

func updateMapInsideLoop() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	for k, v := range m {
		m[2*k] = 2 * v
	}

	fmt.Println(m)
}
