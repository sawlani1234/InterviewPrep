package graph


var root []int 
var size []int

func weightedUnion(x,y int) { 

	rootX,rootY := getRoot(x),getRoot(y)
	if size[rootX] >  size[rootY] {
       size[rootX] += size[rootY]
	   root[rootY] = root[rootX]
	} else {
		size[rootY] += size[rootX]
	   	root[rootX] = root[rootY]
	}
	
}

func getRoot(x int) int {

	for ; x != root[x] ; {
		root[x] = root[root[x]]
		x = root[x]
	}

    return x
}



func Initialise(nodes int) {
	root = make([]int,nodes+1)
	size = make([]int,nodes+1)
	
	for i:=1;i<=nodes;i++ {
		root[i] = i
		size[i] = 1
	}
}












