//time: O(N)
//space: O(N)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sumRange := 0
    var helper func(node *TreeNode)
    helper = func(node *TreeNode) {
        if node == nil { return }
        var sendLeft, sendRight bool
        if node.Val >= low || node.Val <= high {
            if node.Val >= low && node.Val <= high {
                sumRange += node.Val
            }
            sendLeft = node.Val >= low
            sendRight = node.Val <= high
        }
        if sendLeft {
            helper(node.Left)
        }
        if sendRight {
            helper(node.Right)
        } 
    }
    helper(root)
    return sumRange
}