/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 337. House Robber III
// time: O(n)
// space: O(h) ≈ O(n)

func max(a, b int) int {if a >= b { return a}; return b}

func rob(root *TreeNode) int {
    var dfs func(*TreeNode) [2]int
    // [0] -> Y rob this house
    // [1] -> N rob this house; Y rob best child options
    dfs = func(node *TreeNode) [2]int {
        if node == nil {
            return [2]int{0, 0}
        }
        lc := dfs(node.Left)
        rc := dfs(node.Right)
        var ret [2]int
        ret[0] = node.Val + lc[1] + rc[1]
        ret[1] = max(lc[0], lc[1]) + max(rc[0], rc[1])
        return ret
    }
    cand := dfs(root)
    return max(cand[0], cand[1])
}