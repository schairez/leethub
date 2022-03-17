/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    bestSoFar := root.Val
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        nodeVal := float64(node.Val)
        cand := abs(nodeVal - target) 
        if cand < abs(float64(bestSoFar) - target) {
            bestSoFar = node.Val
        }
        if nodeVal > target {
            dfs(node.Left)
        } else {
            dfs(node.Right)
        }
    }
    dfs(root)
    return bestSoFar
}


func max(a, b float64) float64 {
    if a >= b { return a}
    return b
}

func abs(a float64) float64 {
    if a <= 0 {
        return -a
    }
    return a
}



