package graph

import (
	"container/list"
	"fmt"
)

func Bfs(graph [][]int) {
	visited := make([]bool, len(graph))
	l := list.New()
	l.PushBack(1)
	visited[1] = true

	for l.Len() > 0 {
		node := l.Front().Value.(int)
		fmt.Println(node)
		l.Remove(l.Front())
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				l.PushBack(neighbor)
			}
		}
	}
}