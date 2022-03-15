/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {if a >= b { return a}; return b}

func maxPathSum(root *TreeNode) int {
    res := -1 << 32
    dfs(root, &res)
    return res
}




// dfs
func dfs(node *TreeNode, res *int) int {
    if node == nil {
        return 0
    }
    pathSum := node.Val
    leftPath := dfs(node.Left, res)
    rightPath := dfs(node.Right, res)
    if leftPath > 0 {
        pathSum += leftPath
    }
    if rightPath > 0 {
        pathSum += rightPath 
    }
    if pathSum > *res {
        *res = pathSum
    }
    pathSum = node.Val
    pathSum += max(max(leftPath, rightPath), 0)
    return pathSum
}

