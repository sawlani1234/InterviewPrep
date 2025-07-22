package graph

import "fmt"



func recur(node int,graph [][]int,visited []bool) {
	visited[node] = true
	fmt.Println(node)
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			recur(neighbor,graph,visited)
		}
	}
}


func Dfs(graph [][]int) {
	visited := make([]bool, len(graph))	
	recur(1,graph,visited)
}