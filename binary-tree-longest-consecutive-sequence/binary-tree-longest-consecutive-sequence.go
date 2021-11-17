/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
    max := func(a, b int) int {
        if a >= b { return a }
        return b
    }
    res := 1
    
    var dfs func(node *TreeNode, prevVal, sumSoFar int)  
    dfs = func(node *TreeNode, prevVal, sumSoFar int)  {
        if node == nil {
            return
        }
        if prevVal+1 == node.Val {
            sumSoFar += 1
            res = max(res, sumSoFar)
        } else {
            sumSoFar = 1
        }
        dfs(node.Left, node.Val, sumSoFar)
        dfs(node.Right, node.Val, sumSoFar)
    }
    dfs(root, root.Val, 0)
    return res
}