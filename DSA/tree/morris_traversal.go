package tree


func MorrisInorderTraversal(A * TreeNode)  ([]int) {
    // morris bitches
    inorder := []int{}
    
    for ;A!=nil; {
        
        if A.Left == nil {
            inorder = append(inorder,A.Val)
            A = A.Right
        } else {
           temp := A.Left
           
           for ;temp.Right != nil && temp.Right != A; {
               temp = temp.Right
           }
           
           if temp.Right == nil {
               temp.Right = A 
               A = A.Left 
           } else {
               inorder = append(inorder,A.Val)
               temp.Right = nil 
               A = A.Right 
           }
        }
    }
    
    return inorder
}
