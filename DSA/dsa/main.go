package main

import (
	"fmt"
	"solid_design/dsa/graph"
	"solid_design/dsa/sorting"
	// "solid_design/dsa/strings"
	//"solid_design/dsa/tree"
)

func Sort(a []int, strategy sorting.SortingStrategy) []int {
	s := sorting.Sorter{}
	s.SetSortingStrategy(strategy)
	return s.Strategy.Sort(a)

}

func getEdges() [][]int { 
	return [][]int {
		{1,2},
		{1,3},
		{2,4},
		{2,5},
		{3,6},
		{4,5},
		{5,6},
	}
}

/**
lets say graph be like 


		1
	    /\
	    2 3
	   /\ \
      4-5-6


**/

func getWeightedEdges() []graph.EdgeWeighted { 

	return []graph.EdgeWeighted {
		graph.NewEdgeWeighted(1,2,4),
		graph.NewEdgeWeighted(1,3,8),
		graph.NewEdgeWeighted(2, 4,  1),
		graph.NewEdgeWeighted(2, 5, 5),
		graph.NewEdgeWeighted(3, 6, 4),
		graph.NewEdgeWeighted(4, 5, 9),
		graph.NewEdgeWeighted(5, 6, 1),
	}
}

func getDictionary() []string {

	return []string{"shubham","nishi","cheshta","shweta","poorvy"}
}

/*
	A B A B C A B A B
    0 0 1 2
*/
func getLps(pattern string) []int {
	i,j := 0,1
	lps := make([]int,len(pattern))

	for ;j<len(pattern); {

		if pattern[i] == pattern[j] {
			i++
			lps[j] = i
			j++
		} else {
			if i==0 {
				j++
				continue
			} 
			i = lps[i-1]
		}
	}

	return lps
}


type Heap struct {
	arr []int
}

/*
        9
	    /\
	   8  7
	  /\   /\
	 6 5  4  3 
	 
	 

	    5
	   /  \
	  6    4	
      /\    /\
     9	7   8 -1


	     9
		5  2
	10  6   	 
*/

func NewHeap() *Heap  {

    newArr := make([]int,0)

	return &Heap{arr : newArr}
}

func (h *Heap) Len() int {
	return len(h.arr)
}


/*
      2
      /\
	 3  5
	/\   /\
	4 8 9 6
	/\  /
   10 7 -1




*/

func (h *Heap) Heapify(idx int) {

	 if idx == -1 {
		return 
	 }

	 left := 2*idx+1
	 right := 2*idx+2

	 smallest := idx

	if left < h.Len() && h.arr[smallest] >  h.arr[left] {
		smallest = left
	} 
	
	if right < h.Len() && h.arr[smallest] > h.arr[right] {
		smallest = right
	}

	if smallest != idx {
		h.arr[idx],h.arr[smallest] = h.arr[smallest],h.arr[idx]
		h.Heapify(smallest)

	}

 
	
}


func (h *Heap) Push(x int) {     
	h.arr = append(h.arr, x)

	idx := h.Len()-1
	for idx > 0 && h.arr[(idx-1)/2] >  h.arr[idx] {
		h.arr[(idx-1)/2] , h.arr[idx] = h.arr[idx],h.arr[(idx-1)/2]
		idx = (idx-1)/2
	}
	//h.Heapify(h.Len()/2-1)
}

func (h *Heap) Pop() int {
	
	  if h.Len()  == 0 {
		return -1
	  }

	  x := h.arr[0]
	  h.arr[0] , h.arr[h.Len()-1]  = h.arr[h.Len()-1],h.arr[0]
	  h.arr = h.arr[:len(h.arr)-1]
	  h.Heapify(0)

	  return x

}






func main() {
	// graph.Djisktra(graph.GetWeightedGraph(getWeightedEdges(),6),6)
    // trieTest()
    
	arr := []int{10,9,8,7,6,5,4,3,2,1}

	h := NewHeap()

	for _,val := range arr {
		h.Push(val)
	}
	
	fmt.Println(h.arr)
	fmt.Println("-------")
	for i:=0;i<len(arr);i++ {
		fmt.Println(h.Pop())
	}

}


func recur(idx ,k, A int) [][]int {
    
    if k == 0 {
        return make([][]int,0)
    }
    ans := make([][]int,0)

    for i:=idx;i<=A;i++ {
        temp := recur(idx+1,k-1,A)
        for j:=0;j<len(temp);j++ {
			ans = append(ans, append([]int{idx}, temp[j]...))
        }
    }
    
    return ans
}
 
func combine(A int , B int )  ([][]int) {
    // q := list.New()
    // return combine(1,B,A)
    return [][]int{}
}
