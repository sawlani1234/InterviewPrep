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

func PostOrderTraversalRecursion(node *TreeNode) {
	if node == nil {
		return
	}

	PostOrderTraversalRecursion(node.Left)
	PostOrderTraversalRecursion(node.Right)
	fmt.Println(node.Val)
}

func PostOrderTraversalStack(node *TreeNode) {
	l := list.New()

	pushLeft(l, node)
	for l.Len() > 0 {
		temp := l.Back().Value.(*TreeNode)
		if temp.Right != nil {
			pushLeft(l, temp.Right)
		}

	}

}
