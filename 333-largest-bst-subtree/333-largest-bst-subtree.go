/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int { if a >= b { return a}; return b}
func min(a, b int) int { if a <= b { return a}; return b}


func largestBSTSubtree(root *TreeNode) int {
    largestBST := 0
    
    type nodeData struct {loBound, hiBound, numNodes int }
    var dfs func(node *TreeNode) nodeData
    
    dfs = func(node *TreeNode) nodeData {
        if node == nil {
            return nodeData{
                loBound: 1 << 31 - 1,
                hiBound: -1 << 31,
                numNodes: 0,
            }
        }
        fromL := dfs(node.Left)
        fromR := dfs(node.Right)
        data := nodeData{
            loBound: min(node.Val, fromL.loBound),
            hiBound: max(node.Val, fromR.hiBound),
            numNodes: -1,
        }
        if fromL.numNodes != -1 && fromR.numNodes != -1 &&
        node.Val > fromL.hiBound && node.Val < fromR.loBound {
            numNodes := 1 + fromL.numNodes + fromR.numNodes
            largestBST = max(largestBST, numNodes)
            data.numNodes = numNodes
        } 
        
        return data
    }
    dfs(root)
    return largestBST
    
}











