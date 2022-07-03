/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
    n := len(nums)
    stack := make([]*TreeNode, 0, n)
    // 1 <= nums.length <= 1000
    rootNode := &TreeNode{Val: nums[0]}
    stack = append(stack, rootNode) // mono decr
    for i := 1; i < n; i++ {
        currNode := &TreeNode{Val: nums[i]}
        if currNode.Val > rootNode.Val {
            rootNode = currNode
        }
        var leftChild *TreeNode
        for {
            if len(stack) == 0 || currNode.Val < stack[len(stack)-1].Val {
                break
            }
            leftChild, stack = stack[len(stack)-1], stack[:len(stack)-1]
        }
        currNode.Left = leftChild
        if len(stack) != 0 {
            topNode := stack[len(stack)-1]
            topNode.Right = currNode
        }
        stack = append(stack, currNode)
        
    }
    return rootNode
}


