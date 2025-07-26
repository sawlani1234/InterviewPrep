package tree

import (
	"container/list"
	"fmt"
)

/*
		1
      /  \
     2    3
    /\    /
  4   5  6
 /\    \
8 9    10

*/

type PreOrderTraversal struct {
}

func (p *PreOrderTraversal) TraversalRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, ",")
	p.TraversalRecursion(node.Left)
	p.TraversalRecursion(node.Right)
}

func (p *PreOrderTraversal) TraversalLoop(node *TreeNode) {
	l := list.New()
	l.PushBack(node)

	for l.Len() > 0 {
		temp := l.Back()
		l.Remove(l.Back())
		fmt.Print(temp.Value.(*TreeNode).Val, ",")

		if temp.Value.(*TreeNode).Right != nil {
			l.PushBack(temp.Value.(*TreeNode).Right)
		}
		if temp.Value.(*TreeNode).Left != nil {
			l.PushBack(temp.Value.(*TreeNode).Left)
		}
	}

}
