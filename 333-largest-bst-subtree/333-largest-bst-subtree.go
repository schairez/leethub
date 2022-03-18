/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(a, b int) int { if a <= b { return a}; return b}
func max(a, b int) int { if a >= b { return a}; return b}



// interval = [leftI, e]

func largestBSTSubtree(root *TreeNode) int {
    maxBSTSize := 0
    
    type nodeData struct {leftI, rightI, numBSTNodes int}
    var dfs func(*TreeNode) nodeData 
    
    dfs = func(node *TreeNode) nodeData {
        if node == nil {
            return nodeData{leftI: 1 << 31 - 1, 
                            rightI: -1 << 31, 
                            numBSTNodes: 0 }
        }
        data := nodeData{numBSTNodes: -1}
        fromL := dfs(node.Left)
        fromR := dfs(node.Right)
        
        validBSTNodeCond := fromL.numBSTNodes != -1 && fromR.numBSTNodes != -1 &&
                        node.Val > fromL.rightI && node.Val < fromR.leftI 
        
        if validBSTNodeCond {
            data.numBSTNodes = 1 + fromL.numBSTNodes + fromR.numBSTNodes
            maxBSTSize = max(maxBSTSize, data.numBSTNodes)
        }
        data.leftI = min(fromL.leftI, node.Val)
        data.rightI = max(fromR.rightI, node.Val)
        return data
    }
    
    dfs(root)
    return maxBSTSize
    
}


/*
func dfs(node *TreeNode) NodeData {
    if node == nil {
        return NodeData{leftI: 1 << 31 - 1,
                        rightI: -1 << 32,
                        numBSTNodes: 0
                       }
    }
    fromL := dfs(node.Left)
    fromR := dfs(node.Right)
    nodeData := NodeData{}
    if node.Val 
    
    
}
*/