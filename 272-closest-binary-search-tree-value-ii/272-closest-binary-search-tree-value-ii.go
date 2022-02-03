/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(val float64) float64 {if val < 0 {return -val}; return val }


func closestKValues(root *TreeNode, target float64, k int) []int {
    res := make([]int, 0, k)
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode)  {
        if node == nil {
            return
        }
        inorder(node.Left)
        if len(res) < k {
            res = append(res, node.Val)
        } else if abs(float64(node.Val) - target) < abs(float64(res[0])-target) {
            res = res[1:]
            res = append(res, node.Val)
        } else {
            return
        }
        inorder(node.Right)
    }
    inorder(root) 
    return res
}
