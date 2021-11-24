/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
time: O(n)
space: O(1)
*/

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    var successor *TreeNode
    for root != nil  {
        switch {
        case root.Val <= p.Val:
            root = root.Right
        default: //p.Val < root.Val
            //found = true
            successor = root
            root = root.Left
        } 
    }
    return successor
}