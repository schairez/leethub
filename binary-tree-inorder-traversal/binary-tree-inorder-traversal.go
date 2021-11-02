/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//inorder -> LDR

//iterative
//time:O(n); space:O(n)

func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, root.Val)
        root = root.Right
    }
    return res
}




//trivial recursive v
//space:O(n)
//time:O(n)
func inorderTraversalR(root *TreeNode) []int {
    if root == nil { return []int{} }
    res := []int{}
    var helper func(node *TreeNode) 
    helper = func(node *TreeNode) {
        if node == nil { return }
        //LDR; inorder
        helper(node.Left)
        res = append(res, node.Val)
        helper(node.Right)
    } 
    helper(root)
    
    return res
}
