package mapv2

import (
	"fmt"
	"testing"
	"time"
)

func Test_mapconcurrent(t *testing.T) {
	const MAX = 100
	m := NewMap()
	newMap := make(map[int]int)
	for i := 0; i < MAX; i++ {
		newMap[i] = i % 7
	}

	m.Load(newMap)
	fmt.Println(m.Get(3))

	// make 2000 goroutines to read map
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Duration(i%MAX)*time.Millisecond + time.Duration(i/10)*time.Second)
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
