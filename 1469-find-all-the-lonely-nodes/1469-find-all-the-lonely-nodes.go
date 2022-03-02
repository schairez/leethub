/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) []int {
    var (
        node *TreeNode
        res []int
    )
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        node, queue = queue[0], queue[1:]
        if node.Left == nil {
            if node.Right != nil {
                queue = append(queue, node.Right)
                res = append(res, node.Right.Val)
            }
        } else if node.Right == nil {
            if node.Left != nil {
                queue = append(queue, node.Left)
                res = append(res, node.Left.Val)
            }
        } else {
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
            
        }
    }
    
    return res
}