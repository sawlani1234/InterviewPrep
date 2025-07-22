package graph

import (
	"container/heap"
	"fmt"
)

// import "container/heap"


type WeightedEdge []EdgeWeighted

func (e *WeightedEdge) Len() int {
	return len(*e)
}

func (e *WeightedEdge) Push(x any) {
   *e = append(*e, x.(EdgeWeighted))
}

func (e *WeightedEdge) Pop() (y any)  {
	*e, y = (*e)[:e.Len()-1], (*e)[e.Len()-1]
	return
}

func (e *WeightedEdge) Less(i,j int) bool {
	return (*e)[i].weight <  (*e)[j].weight
}

func (e *WeightedEdge) Swap(i,j int) {
	(*e)[i] , (*e)[j] = (*e)[j], (*e)[i]
}


func Prims(graph [][]EdgeWeighted,nodes int) {
     
	h := &WeightedEdge{}
	heap.Init(h)
    heap.Push(h,EdgeWeighted{x: -1,y:1,weight: 0})
	vis := make([]bool,nodes+1)

	for ;h.Len() > 0 ; {

		x := heap.Pop(h).(EdgeWeighted)
		if vis[x.y] {
			continue
		}
		vis[x.y] = true 
		fmt.Println("Node :",x.y, " Weight :",x.weight)

		for i:=0;i<len(graph[x.y]);i++ {
			
			n := graph[x.y][i]
			
			if !vis[n.y] {
				heap.Push(h,EdgeWeighted{x:-1,y:n.y,weight: n.weight})
			}

		}
		
	}

}

