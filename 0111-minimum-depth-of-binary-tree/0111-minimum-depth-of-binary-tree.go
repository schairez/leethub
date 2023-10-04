/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
// time: O(n)
// space: O(n)

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxInt32 := 1 << 31 - 1
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return maxInt32 }
        minFromLeft := dfs(node.Left)
        minFromRight := dfs(node.Right)
        if minFromLeft == maxInt32 && minFromRight == maxInt32 {
            return 1
        }
        return 1 + min(minFromLeft, minFromRight)
    }
    return dfs(root)
}


func min(a, b int) int { if a <= b { return a}; return b}