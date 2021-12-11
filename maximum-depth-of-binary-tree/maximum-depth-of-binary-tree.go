/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//postorder traversal
//time: O(n)
//space: O(n)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    fromLeft := maxDepth(root.Left) 
    fromRight := maxDepth(root.Right) 
    if fromLeft >= fromRight {
        return fromLeft+1
    }
    return fromRight+1
}