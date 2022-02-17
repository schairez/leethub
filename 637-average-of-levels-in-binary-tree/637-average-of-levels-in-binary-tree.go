/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func avg(sum float64, n int) float64 {
    return sum / float64(n)
}
//time: O(numNodes)
//space: O(numNodes At maxLvl)

func averageOfLevels(root *TreeNode) []float64 {
    var res []float64
    queue := []*TreeNode{root}
    var node *TreeNode 
    for len(queue) != 0 {
        n := len(queue)
        var sumNodesAtLvl float64
        for i := 0; i < n; i++ {
            node, queue = queue[0], queue[1:]
            sumNodesAtLvl += float64(node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, avg(sumNodesAtLvl, n))
    } 
    return res
}