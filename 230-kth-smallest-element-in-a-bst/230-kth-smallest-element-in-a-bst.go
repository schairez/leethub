/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    kthVal := k
    var res int
    inorder(root, func(val int) bool {
        kthVal--
        if kthVal == 0 {
            res = val
            return true
        }
        return false
    })
    
    return res
}

func inorder(node *TreeNode, checkFn func(val int) bool) {
    if node == nil {
        return
    }
    inorder(node.Left, checkFn)
    if done := checkFn(node.Val); done == true {
        return
    }
    inorder(node.Right, checkFn)
}