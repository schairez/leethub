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
        var sendLeft, sendRight, addToRange bool
        addToRange = node.Val >= low && node.Val <= high 
        sendLeft = node.Val >= low
        sendRight = node.Val <= high
        if addToRange { 
            sumRange += node.Val
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