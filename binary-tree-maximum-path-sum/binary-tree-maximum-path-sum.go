/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    globalMaxPath := -1 << 31
    
    postOrder(root, func(localMaxPath int) {
        globalMaxPath = max(globalMaxPath, localMaxPath)
    })
    
    return globalMaxPath
}

    
//post -> LRD
func postOrder(node *TreeNode, f func(localMaxPath int)) int {
    if node == nil {
        return 0
    }
    leftGain := max(postOrder(node.Left, f), 0) 
    rightGain := max(postOrder(node.Right, f), 0) 
    //contribution of nodeVal and its l,r children
    maxPathWChildren := node.Val + leftGain + rightGain
    
    f(maxPathWChildren)
    
    return node.Val + max(leftGain, rightGain)
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

