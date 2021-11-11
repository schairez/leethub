/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil { return false }
    if (root.Val - targetSum == 0) && root.Left == nil && root.Right == nil { 
        return true 
    }
    return hasPathSum(root.Left, targetSum-root.Val) || 
    hasPathSum(root.Right, targetSum-root.Val) 
  
    
}