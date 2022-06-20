/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 951. Flip Equivalent Binary Trees
// time: O(min(n1, n2))
// space: O(min(h1, h2))

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil || root2 == nil {
        if root1 == nil && root2 == nil {
            return true
        }
        return false
    }
    if root1.Val != root2.Val {
        return false
    }
    return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) ||
    flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}