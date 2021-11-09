//time: O(n)
//space: O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func sumOfLeftLeaves(root *TreeNode) int {
    res := 0 //sumLeftLeaves
    stack := []*TreeNode{}    
    
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
            if root == nil { break }
            if root.Left == nil && root.Right == nil {
                res += root.Val
            }
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = root.Right
    } 
    
    return res
}