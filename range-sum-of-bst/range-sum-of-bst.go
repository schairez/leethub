/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//time: O(n)
//space:O(n)
//iterative inorder LDR
func rangeSumBST(root *TreeNode, low int, high int) int {
    res := 0
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        addToSumCond := root.Val >= low && root.Val <= high
        if addToSumCond { res += root.Val }
        root = root.Right
    }
    return res
}


//time: O(N)
//space: O(N)
//recursive DLR preorder
func rangeSumBSTRecursive(root *TreeNode, low int, high int) int {
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