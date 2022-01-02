/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//if balanced tree height is logn = H
//recursive call stack time and space avg case: O(log(n))
//worst case: O(n)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        newNode := &TreeNode{Val: val}
        return newNode
    }
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    } else {
        root.Left = insertIntoBST(root.Left, val)
    }
    return root
}
