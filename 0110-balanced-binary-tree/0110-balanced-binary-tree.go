/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//time: O(n)
//space: O(n)

func isBalanced(root *TreeNode) bool {
    return dfsHeight(root) != -1   
}


func dfsHeight(node *TreeNode) int {
    if node == nil { return 0}
    leftHeight := dfsHeight(node.Left)
    if leftHeight == -1 { return -1}
    rightHeight := dfsHeight(node.Right)
    if rightHeight == -1 { return -1}

    if abs(leftHeight - rightHeight) <= 1 {
        return 1 + max(leftHeight, rightHeight)   
    }
    return -1
}
func max(a, b int) int {
    if a >= b { return a}
    return b
}

func abs(a int) int {
    if a < 0 { return -a}
    return a
}