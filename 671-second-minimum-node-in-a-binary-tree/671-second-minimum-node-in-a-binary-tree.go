/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//time: O(n)
//space: O(heightTree) theta(logn) ; worst case O(n) where n = numNodes

func findSecondMinimumValue(root *TreeNode) int {
    if root.Left == nil && root.Right == nil {
        return -1
    }
    maxInt32 := (1 << 31) -1
    res := maxInt32
    foundRes := false
    minNode := root.Val
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Val == minNode {
            dfs(node.Left)
            dfs(node.Right)
        } else {
            if node.Val > minNode && node.Val <= res {
                foundRes = true
                res = node.Val
            }
        }
    }
    dfs(root)
    if !foundRes {
        return -1
    }
    return res
    
}