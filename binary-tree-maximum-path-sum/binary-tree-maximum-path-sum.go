/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(vals ...int) int {
    maxV := vals[0]
    for i:=1; i<len(vals);i++ {
        if vals[i] > maxV {
            maxV = vals[i]
        }
    }
    return maxV
}
    
func maxPathSum(root *TreeNode) int {
    globalMaxPath := -1 << 31
    
    postOrder(root, func(localMaxPath int) {
        globalMaxPath = max(globalMaxPath, localMaxPath)
    })
    return globalMaxPath
}

//post -> LRD
func postOrder(node *TreeNode, f func(localMaxPath int)) int {
    if node == nil {
        return 0
    }
    leftGain := max(postOrder(node.Left, f), 0) 
    rightGain := max(postOrder(node.Right, f), 0) 
    //contribution of nodeVal and its l,r children
    maxPathWChildren := node.Val + leftGain + rightGain
    f(maxPathWChildren)
    
    return node.Val + max(leftGain, rightGain)
}



/*

func inOrder(curNode *TreeNode, f func(curNode *TreeNode)) {
    if curNode == nil {
        return
    }
    inOrder(curNode.Left, f)
    f(curNode)
    inOrder(curNode.Right, f)
}


func maxPathSum(root *TreeNode) int {
    
    max := func(a, b int) int {
        if a >= b {
            return a
        }
        return b
    }
    globalMaxPath := -1 << 31
    localMaxPath := globalMaxPath
    
    inOrder(root, func(curNode *TreeNode) {
        nodeV := curNode.Val
        localMaxPath = max(localMaxPath + nodeV, nodeV)
        globalMaxPath = max(globalMaxPath, localMaxPath)
    })
    
    return globalMaxPath
}

*/


