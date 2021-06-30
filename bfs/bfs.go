package bfs

import "fmt"

func InitGraph() [][]bool {
	result := make([][]bool, 10)
	result[0] = []bool{false,true,true,false,false,false,false,false,true,false}
	result[1] = []bool{true,false,true,false,false,true,false,false,false,false}
	result[2] = []bool{true,false,false,true,false,true,false,false,false,false}
	result[3] = []bool{false,false,true,false,true,false,false,true,false,true}
	result[4] = []bool{false,false,false,true,false,false,false,false,true,false}
	result[5] = []bool{false,true,true,false,false,false,false,false,false,false}
	result[6] = []bool{false,false,false,false,false,false,false,false,false,true}
	result[7] = []bool{false,false,false,true,false,false,false,false,false,false}
	result[8] = []bool{true,false,false,false,true,false,false,false,false,false}
	result[9] = []bool{false,false,false,true,false,false,true,false,false,false}
	return result
}

func BFS(index int, N int)  {
	index--
	arr := InitGraph()

	checked := make([]bool, N)
	checked[index] = true

	queue := make([]int, 0)
	queue = append(queue, index)
	for {
		if len(queue) == 0 {
			fmt.Println("finish BFS with vertex ", index+1)
			return
		}

		vertex := queue[0]
		queue = queue[1:]
		fmt.Println("Visit vertex ", vertex+1)
		for i := 0; i < N; i++ {
			if arr[vertex][i] && !checked[i] {
				queue = append(queue, i)
				checked[i] = true
			}
		}
	}
}

