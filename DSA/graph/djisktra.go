package graph

import (
	"container/heap"
)

type Pair struct {
	v int
	weight int 
}

type DjisktraHeap []Pair


func (d *DjisktraHeap) Len() int {
	return len(*d)
}

func (d *DjisktraHeap) Less(i,j int) bool {
	return (*d)[i].weight < (*d)[j].weight
}

func (d *DjisktraHeap) Swap(i,j int) {
	(*d)[i],(*d)[j] = (*d)[j],(*d)[i]
}

func (d *DjisktraHeap) Push(x any) {
	(*d) = append((*d), x.(Pair))
}

func (d *DjisktraHeap) Pop() (x any) {
	x,(*d)  = (*d)[d.Len()-1],(*d)[:d.Len()-1]
	return x
}

func Djisktra(gr [][]EdgeWeighted,A int) []int{

	dist := make([]int,A+1)
	vis := make([]bool,A+1)

	for i:=1;i<=A;i++ {
		dist[i] = 10000000
	}

	dist[1] = 0
	pq := &DjisktraHeap{}
	heap.Init(pq)
	pq.Push(Pair{1,0})

	for pq.Len() > 0 {

		ele := heap.Pop(pq).(Pair)

		
		if !vis[ele.v] {
			vis[ele.v] = true 

			for j:=0;j<len(gr[ele.v]);j++ {

				child := gr[ele.v][j]
				
				if !vis[child.y] && dist[ele.v]+ child.weight < dist[child.y] {
					dist[child.y] = dist[ele.v]+ child.weight
					heap.Push(pq,Pair{child.y,dist[child.y]})
				}
			}
		}
	}
	return dist
}