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

type InorderTraversal struct {
}

func (i *InorderTraversal) TraversalRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	i.TraversalRecursion(node.Left)
	fmt.Print(node.Val, ",")
	i.TraversalRecursion(node.Right)
}

func pushLeft(l *list.List, node *TreeNode) {
	for node != nil {
		l.PushBack(node)
		node = node.Left
	}
}

func (i *InorderTraversal) TraversalLoop(node *TreeNode) {
	l := list.New()

	pushLeft(l, node)

	for l.Len() > 0 {
		fmt.Print(l.Back().Value.(*TreeNode).Val, ",")
		temp := l.Back().Value.(*TreeNode).Right
		l.Remove(l.Back())
		if temp != nil {
			pushLeft(l, temp)
		}
	}

}
