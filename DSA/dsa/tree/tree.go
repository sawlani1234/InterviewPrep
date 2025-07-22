package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

/*
		1
      /  \
     2    3
    /\    /
  4   5  6
 /\    \
8 9    10

*/

var sampleTree = []int{1, 2, 3, 4, 5, 6, -1, 8, 9, -1, 10, -1, -1}
var Root *TreeNode

func init() {
	Root = generateTreeFromArr()
}

type NodePair struct {
	node *TreeNode
	idx  int
}

func generateTreeFromArr() *TreeNode {
	root := NewTreeNode(sampleTree[0])

	cur := root

	q := list.New()
	q.PushBack(NodePair{node: cur, idx: 0})

	for q.Len() > 0 {
		temp, idx := q.Front().Value.(NodePair).node, q.Front().Value.(NodePair).idx

		if (2*idx+1 < len(sampleTree)) && (sampleTree[2*idx+1] != -1) {
			temp.Left = NewTreeNode(sampleTree[2*idx+1])
			q.PushBack(NodePair{node: temp.Left, idx: 2*idx + 1})
		}

		if 2*idx+2 < len(sampleTree) && (sampleTree[2*idx+2] != -1) {
			temp.Right = NewTreeNode(sampleTree[2*idx+2])
			q.PushBack(NodePair{node: temp.Right, idx: 2*idx + 2})

		}
		q.Remove(q.Front())
	}

	return root
}

type TraversalStrategyStruct struct {
	Trav TraversalStrategyInterface
}

func NewTraversalStrategyStruct() *TraversalStrategyStruct {
	return &TraversalStrategyStruct{}
}

func (t *TraversalStrategyStruct) SetStrategy(traversal TraversalStrategyInterface) {
	t.Trav = traversal
}

type TraversalStrategyInterface interface {
	TraversalRecursion(node *TreeNode)
	TraversalLoop(node *TreeNode)
}
