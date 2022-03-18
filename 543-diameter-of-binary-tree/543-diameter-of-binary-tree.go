/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int { if a >= b { return a}; return b }


func diameterOfBinaryTree(root *TreeNode) int {
    longestPath := 0
    var dfs func(*TreeNode) int
    
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        fromL := dfs(node.Left) 
        fromR := dfs(node.Right) 
        cand := fromL + fromR 
        longestPath = max(longestPath, cand)
        return max(fromL, fromR) + 1
    }
    
    dfs(root)
    
    return longestPath
}