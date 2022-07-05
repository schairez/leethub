/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 865. Smallest Subtree with all the Deepest Nodes
// time: O(n)
// space: O(h) ≈ O(n)

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    type pair struct {depth int; node *TreeNode}
    
    var dfs func(node *TreeNode) pair
    dfs = func(node *TreeNode) pair {
        if node == nil {
            return pair{0, nil}
        }
        lc := dfs(node.Left)
        rc := dfs(node.Right)
        if lc.depth == rc.depth {
            return pair{1 + lc.depth, node}
        } else if lc.depth > rc.depth {
            return pair{ 1 + lc.depth, lc.node}
        }
        return pair { 1 + rc.depth, rc.node}
    }
    
    res := dfs(root)
    return res.node
}