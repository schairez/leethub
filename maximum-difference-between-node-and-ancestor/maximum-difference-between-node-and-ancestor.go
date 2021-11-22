/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    resMaxV := -1*(1 << 31)
    postOrder(root, func(minV, maxV, nodeV int) {
        lMax := abs(nodeV - minV)
        lMax = max(lMax, abs(nodeV - maxV))
        resMaxV = max(lMax, resMaxV)
    })
    return resMaxV
}

func abs(a int) int { if a < 0 { return -a }; return a }
func min(a, b int) int { if a <= b { return a}; return b}
func max(a, b int) int { if a >= b { return a}; return b}

func postOrder(node *TreeNode, f func(minV, maxV, nodeV int)) (int,int) {
    if node == nil {
        maxV := -1*(1 << 31)
        minV := (1 << 31)-1
        return minV, maxV
    }
    nLMin, nLMax := postOrder(node.Left, f)
    nRMin, nRMax := postOrder(node.Right, f)
    nodeV := node.Val
    minV := min(nLMin, nRMin)
    minV = min(minV, node.Val)
    maxV := max(nRMax, nLMax)
    maxV = max(maxV, node.Val)
    f(minV, maxV, nodeV)
    return minV, maxV
}