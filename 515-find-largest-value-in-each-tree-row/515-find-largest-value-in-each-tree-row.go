/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0)
    var dfs func(node *TreeNode, idx int)
    dfs = func(node *TreeNode, idx int) {
        if node == nil {
            return
        }
        if len(res) == 0 || idx == len(res) {
            res = append(res, node.Val)
        } else {
            if res[idx] < node.Val {
                res[idx] = node.Val
            }
        }
        dfs(node.Left, idx+1)
        dfs(node.Right, idx+1)
    }
    dfs(root, 0)
    return res
    
}