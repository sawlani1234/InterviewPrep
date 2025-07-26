package graph 



/**
lets say graph be like 


		1
	    /\
	    2 3
	   /\ \
      4-5-6


**/

type EdgeWeighted struct {
	x int
	y int
	weight int
}

func NewEdgeWeighted(x,y,weight int) EdgeWeighted {
	return EdgeWeighted{x:x,y:y,weight: weight}
}

func GetWeightedGraph(edges []EdgeWeighted,nodes int) [][]EdgeWeighted {
     graph := make([][]EdgeWeighted,nodes+1)

	 for _,edge := range edges {
		u,v,w := edge.x,edge.y,edge.weight
		graph[u] = append(graph[u], EdgeWeighted{y:v,weight: w})
		graph[v] = append(graph[v], EdgeWeighted{y:u,weight: w})
	 }
	 return graph
}


func GetGraph(edges [][]int,nodes int) [][]int {
	graph := make([][]int,nodes+1)
	for _,edge := range edges {
		u,v := edge[0],edge[1]
		graph[u] = append(graph[u],v)
		graph[v] = append(graph[v],u)
	}
	return graph
}
