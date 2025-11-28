package tree 

var Tree []int

func Build(nums []int,l,r,node int) {

	if l == r {
		Tree[node] = nums[l]
		return 
	} 

	mid := (l+r)>>1
	Build(nums,l,mid,2*node)
	Build(nums,mid+1,r,2*node+1)

	Tree[node] = Tree[2*node]+Tree[2*node+1]
	

}



func Query(start,end,l,r,node int)  int { 

	if start > r || end < l {
		return 0
	}

	if  start >=l && end<=r {
		return Tree[node]
	}
	mid := (start+end)>>1
	
	return Query(start,mid,l,r,2*node)+Query(mid+1,end,l,r,2*node+1)
	
}


