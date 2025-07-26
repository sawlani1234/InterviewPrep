package graph

var topo []int
func DfsTopo(graph [][]int,cur int ,vis []bool) {
    
    for i:=0;i<len(graph[cur]);i++ {
        node := graph[cur][i]
        if !vis[node] {
            vis[node] = true 
            DfsTopo(graph,node,vis)
        }
    }
    
    topo = append(topo,cur)

	// adddd for check
    
} 