/*
time: O(n)
space: O(n)
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    //precondition: numNodes in [1,10^4]
    stack := []*TreeNode{}
    var pre *TreeNode
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pre != nil && pre.Val  >= root.Val {
            return false
        }
        pre = root
        root = root.Right
    }
    return true
    
}




/*
root = [5,1,4,null,null,3,6]
                5            #vals to right are > parent for bst validity
               / \ 
              1  4
                 /\
                3  6
           #append 
           st -> [5] 
           st -> [5,1]
           #pop
           r -> st.pop()-> 1; st=[5]
           #assignment
           pre -> r -> 1
           r -> nil
           #pop
           r -> st.pop() -> 5 ; st =[]
           
           
           #assignment
           pre -> r -> 5
           r -> r.right -> 4
           
           
           #append
           st -> [4]; -> [4,3]
           r -> st.pop() -> 3
           
           *****check
           if r != nil && pre.Val >= r.Val { -> F}
           -----> return False
           
           
           
           

*/
