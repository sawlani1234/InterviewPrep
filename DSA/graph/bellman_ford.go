package graph

import "fmt"

// Single Source SHortest Path. O(VE) , detects negative cyle as well
// Based on principle of relaxation 
func BellmanFord(edges [][]int,n int) {

	dist := make([]int,n)
	
	for i:=0;i<n;i++ {
		dist[i] = 1e8
	}
	
	dist[0] = 0

	for i:=0;i<n;i++ {
		for j:=0;j<len(edges);j++ {
			u,v,w := edges[j][0],edges[j][1],edges[j][2]

			if i == (n-1) && dist[u]+w < dist[v] {
				fmt.Println("negative edge cycle")
				return 
			}
			dist[v] = min(dist[v],dist[u]+w)
		}
	}

	fmt.Println(dist)
}