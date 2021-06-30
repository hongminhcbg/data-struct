package dfs

import (
	"data-struct/bfs"
	"fmt"
)

func DFS(index int, N int) {
	arr := bfs.InitGraph()
	index--
	visited := make([]bool, N)

	stack := make([]int, 0)
	stack = append(stack, index)
	for {
		if len(stack) == 0 {
			fmt.Println("finish DFS with vertex ", index+1)
			return
		}

		vertex := stack[0]
		stack = stack[1:]
		if visited[vertex] {
			continue
		}

		visited[vertex] = true
		fmt.Println("Visit vertex ", vertex+1)
		for i := 0; i < N; i++ {
			if arr[vertex][i] && !visited[i] {
				stack = append([]int{i}, stack...)
			}
		}
	}
}
