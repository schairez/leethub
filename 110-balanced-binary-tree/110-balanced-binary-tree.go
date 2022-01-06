//time: O(n) we traverse down all nodes 
//space: O(n) recursive call stack

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil { return true }
    lHeight := getHeight(root.Left)
    rHeight := getHeight(root.Right)
    return abs(lHeight -  rHeight) <= 1 &&
    isBalanced(root.Left) && isBalanced(root.Right)
       
    
}

func abs(a int) int { 
    if a < 0 { return -a }
    return a
}
func max(a, b int) int { if a >= b { return a}; return b}

func getHeight(node *TreeNode) int {
    if node == nil { return 0 }
    
    lHeight := getHeight(node.Left)
    rHeight := getHeight(node.Right)
    
    return max(lHeight, rHeight) + 1
    
}

