package graph

import (
	"fmt"
	"sort"
)



func Kruskal(weightedEdges []EdgeWeighted,nodes int) {

	Initialise(nodes)
	
	sort.Slice(weightedEdges,func(i,j int) bool {
		return weightedEdges[i].weight < weightedEdges[j].weight
	})

	for i:=0;i<len(weightedEdges);i++ {
		if getRoot(weightedEdges[i].x) == getRoot(weightedEdges[i].y) {
			continue
		}	
		fmt.Println(weightedEdges[i].x, " : ",weightedEdges[i].y, " : ", weightedEdges[i].weight)
		weightedUnion(weightedEdges[i].x,weightedEdges[i].y)
	}

}