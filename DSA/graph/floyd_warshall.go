package graph

import (
	"fmt"
)

func FloydWarshall(edges [][]int, n int) {

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)

		for j:=0;j<n;j++ {
			dist[i][j] = 1e7
		}
		dist[i][i] = 0
	}

	for i:=0;i<len(edges);i++ {
		u,v,w := edges[i][0],edges[i][1],edges[i][2]

		dist[u][v] = w 
		dist[v][u] = w
	}

	

	for k:=0;k<n;k++ {
		for i:=0;i<n;i++ {
			for j:=0;j<n;j++ {
				if dist[i][k] != 1e7 && dist[k][j] != 1e7 {
					dist[i][j] = min(dist[i][j],dist[i][k]+dist[k][j])
				}
			}
		}
	}

	for i:=0;i<n;i++ {
		fmt.Println(dist[i])
	}

}