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
    //constraint: -2^31 <= Node.val <= 2^31 - 1
    minInt32 := -1*(1 << 31)
    maxInt32 := (1 << 31) - 1
    var helper func(node *TreeNode, minV, maxV int) bool
    helper = func(node *TreeNode, minV, maxV int) bool {
        if node == nil {
            return true
        }
        cond := node.Val >= minV && node.Val <= maxV
        return cond && 
        helper(node.Left, minV, node.Val-1) && 
        helper(node.Right, node.Val+1, maxV)
    }
    return helper(root, minInt32, maxInt32)
}

//iterative stack
func isValidBSTIterative(root *TreeNode) bool {
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

*/
