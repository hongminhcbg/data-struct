package main

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	fmt.Printf("mem consummed %d bytes\n", m.Alloc)
}

type Foo struct {
	S   string // 8 bytes
	Age uint32 // 4 bytes
}

func memLeak() {
	printAlloc() // init
	m := make(map[string]int)
	for i := 0; i < 1_000_000; i++ {
		m[fmt.Sprint(i)] = i
	}

	printAlloc() // after add 1M records
	for i := 0; i < 1_000_000; i++ {
		delete(m, fmt.Sprint(i))
	}

	runtime.GC()
	printAlloc() // after delete 1M records
	runtime.KeepAlive(m)

	// init: mem consummed 109712 bytes
	// add 1M records: mem consummed 106997136 bytes
	// remove 1M records: mem consummed 58046520 bytes
	// because bucket size of map only decrese
	// Solution: after delete we can lock and recreate map
}
