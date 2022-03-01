/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func getDirections(root *TreeNode, startValue int, destValue int) string {
    var (
        lca func(*TreeNode, int, int) *TreeNode
        dfsPath func(*[]byte, *TreeNode, int) bool 
    )
    
    lca = func(node *TreeNode, srcNodeV, dstNodeV int) *TreeNode {
        if node == nil {
            return nil
        }
        if node.Val == srcNodeV || node.Val == dstNodeV {
            return node
        }
        lNode := lca(node.Left, srcNodeV, dstNodeV)
        rNode := lca(node.Right, srcNodeV, dstNodeV)
        if lNode != nil && rNode != nil {
            return node
        } else if lNode != nil {
            return lNode
        } else if rNode != nil {
            return rNode
        }
        return nil
            
    }
    // dst and src nodes have to pass through lcaNode
    lcaNode := lca(root, startValue, destValue)
    // sb leftBound and sbRightBound
    dfsPath = func(pathBytes *[]byte, node *TreeNode, searchVal int) bool { 
        if node == nil {
            return false
        }
        if node.Val == searchVal {
            return true
        }
        *pathBytes = append(*pathBytes, 'L')
        if foundLeft := dfsPath(pathBytes, node.Left, searchVal); foundLeft == true {
            return true
        }
        n := len(*pathBytes)
        *pathBytes = (*pathBytes)[:n-1]
        *pathBytes = append(*pathBytes, 'R')
        if foundRight := dfsPath(pathBytes, node.Right, searchVal); foundRight == true {
            return true
        }
        n = len(*pathBytes)
        *pathBytes = (*pathBytes)[:n-1]
        return false
    }
    var srcBytes, dstBytes []byte
    dfsPath(&srcBytes, lcaNode, startValue)
    dfsPath(&dstBytes, lcaNode, destValue)
    
    var sb strings.Builder
    sb.Grow(len(srcBytes)+len(dstBytes))
    for idx := range srcBytes {
        srcBytes[idx] = 'U'
        sb.WriteByte(srcBytes[idx])
    }
    for idx := range dstBytes {
        sb.WriteByte(dstBytes[idx])
    }
    return sb.String()
    
}



