/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    var (
        res [][]int
        queue []*TreeNode
        node *TreeNode
    )
    revSlice := func(arr []int) {
        i, j := 0, len(arr)-1
        for i < j {
            arr[i], arr[j] = arr[j], arr[i]
            i++
            j--
        }
    }
    isLeftToRight := true 
    queue = append(queue, root)
    for len(queue) != 0 {
        var nodes []int
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            nodes = append(nodes, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if !isLeftToRight {
            revSlice(nodes)
        }
        res = append(res, nodes)
        isLeftToRight = !isLeftToRight
    }
    
    return res
}