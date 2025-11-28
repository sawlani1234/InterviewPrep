package main

import (
	//"fmt"
	"container/heap"
	"fmt"
	"solid_design/dsa/graph"
	"solid_design/dsa/sorting"
	"solid_design/dsa/tree"
	"sort"
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

/**
		1
	    /\	
	   2  7
          /\ 
	  3 -6 -8
      /\ 
     4. 5



		   */        


func getWeightedGraphExample() [][]graph.EdgeWeighted {
	return graph.GetWeightedGraph([]graph.EdgeWeighted{
		graph.NewEdgeWeighted(7,8,10),
		graph.NewEdgeWeighted(3,5,1),
		graph.NewEdgeWeighted(6,7,2),
		graph.NewEdgeWeighted(1,2,4),
		graph.NewEdgeWeighted(3,6,4),
		graph.NewEdgeWeighted(1,7,6),
		graph.NewEdgeWeighted(3,4,7),
		graph.NewEdgeWeighted(6,8,1),
		},8)
}


type TrieNode struct {
	child [26]*TrieNode
	isEnd bool
}

func NewTrieNode() *TrieNode {
	 var child [26]*TrieNode
	 
	 return &TrieNode{child:child,isEnd:false}
}



func add(root *TrieNode,word string) {
	
	temp := root
	for i:=0;i<len(word);i++ {
		if temp.child[word[i]-'a'] == nil {
			temp.child[word[i]-'a'] = NewTrieNode()
		}
		temp = temp.child[word[i]-'a']
	}

	temp.isEnd = true
}	

func search(root *TrieNode,word string) bool {
	temp := root
	
	for i:=0;i<len(word) && temp != nil;i++ {
		temp = temp.child[word[i]-'a']
	}
	return temp != nil && temp.isEnd
}


func main() {

	tree.Tree = make([]int,40)

	nums := []int{1,2,3,5,6}

	tree.Build(nums,0,4,1)

	fmt.Println(tree.Query(0,4,2,1,1))

    
	// graph.Djisktra(graph.GetWeightedGraph(getWeightedEdges(),6),6)
    // trieTest()

	// graph.BellmanFord([][]int{
	// 	{0,1,6},
	// 	{0,2,5},
	// 	{1,3,1},
	// 	{2,1,-2},
	// 	{3,2,4},
	// },4)

	// graph.FloydWarshall([][]int{
	// 	{0,1,3},
	// 	{0,2,5},
	// 	{1,2,1},
	// 	{1,3,6},
	// 	{2,3,2},
	// },4)
    
	// arr := []int{10,9,8,7,6,5,4,3,2,1}

	// h := NewHeap()

	// for _,val := range arr {
	// 	h.Push(val)
	// }
	
	// fmt.Println(h.arr)
	// fmt.Println("-------")
	// for i:=0;i<len(arr);i++ {
	// 	fmt.Println(h.Pop())
	// }

	
	//fmt.Println(Prims(getWeightedGraphExample(),8))
	// Kruskal([]graph.EdgeWeighted{
	// 	graph.NewEdgeWeighted(7,8,10),
	// 	graph.NewEdgeWeighted(3,5,1),
	// 	graph.NewEdgeWeighted(6,7,2),
	// 	graph.NewEdgeWeighted(1,2,4),
	// 	graph.NewEdgeWeighted(3,6,4),
	// 	graph.NewEdgeWeighted(1,7,6),
	// 	graph.NewEdgeWeighted(3,4,7),
	// 	graph.NewEdgeWeighted(6,8,1)},8)
	// root := NewTrieNode()
	// add(root,"shubham")
	// add(root,"swarnim")
	// add(root,"shubhs")
	// add(root,"shubhs")
	// fmt.Println("sh",search(root,"sh"))
	// fmt.Println("shubhs",search(root,"shubhs"))
	// fmt.Println("swarnim",search(root,"swarnim"))
	// fmt.Println("swar",search(root,"swar"))
	// fmt.Println("shubham",search(root,"shubham"))

}

type IntHeap []Pair 

type Pair struct {
	x int 
	w int 
}



func NewPair(x,w int) Pair {
	return Pair{x,w}
}

func(h *IntHeap) Push(x interface{}){
	(*h) = append((*h), x.(Pair))
}

func(h *IntHeap) Pop() (x interface{}) {
	x , (*h) = (*h)[h.Len()-1] , (*h)[:h.Len()-1]

	return x
}

func(h *IntHeap) Len() int {
	return len(*h)
}

func(h *IntHeap) Swap(i,j int) {
	(*h)[i],(*h)[j] = (*h)[j],(*h)[i]
}

func (h *IntHeap) Less(i,j int) bool {
	return (*h)[i].w < (*h)[j].w
}



func Prims(graph [][]graph.EdgeWeighted,n int) int {

	vis := make([]bool,n+1)

	q := &IntHeap{}
	heap.Init(q)
	heap.Push(q,NewPair(1,0))
	ans := 0

	for ;q.Len() > 0; {

		p := heap.Pop(q).(Pair)
		parent,weight := p.x,p.w
		

		if !vis[parent] {
			vis[parent] = true 
			ans += weight

			for k:=0;k<len(graph[parent]);k++ {
				child := graph[parent][k].GetY()
				w := graph[parent][k].GetWeight()

				if !vis[child] {
					heap.Push(q,NewPair(child,w))
				}
			}
		}
	}

	return ans 

}

var siz []int 
var root []int

func getRoot(a int) int {

	for ;a!=root[a]; {
		root[a] = root[root[a]]
		a = root[a]
	}

	return a
}

func union(a,b int) {

	rootA := getRoot(a)
	rootB := getRoot(b)
	
	if siz[rootA] >= siz[rootB] {
		siz[rootA] += siz[rootB]
		root[rootB] = rootA
	} else {
		siz[rootB] += siz[rootA]
		root[rootA] = root[rootB]
	}
}

func find(a,b int) bool {
	return getRoot(a) != getRoot(b)
}

func Kruskal(graph []graph.EdgeWeighted,n int) {

	ans := 0
   
	siz = make([]int, n+1)
	root = make([]int, n+1)


	for i:=0;i<=n;i++ {
		siz[i] = 1
		root[i] = i 
	}

	sort.Slice(graph,func(i,j int) bool {
		return graph[i].GetWeight() < graph[j].GetWeight()		
	})

	for i:=0;i<len(graph);i++ {
		u,v,w := graph[i].GetX(),graph[i].GetY(),graph[i].GetWeight()

	    if find(u,v) {
			union(u,v)
			ans+=w
		}

	}

	fmt.Println(ans)
}



 
func combine(A int , B int )  ([][]int) {
    // q := list.New()
    // return combine(1,B,A)
    return [][]int{}
}
