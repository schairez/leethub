/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func invertTree(root *TreeNode) *TreeNode {
    if root == nil { return root }
    q := []*TreeNode{}
    q = append(q, root, root.Left, root.Right)
    for len(q) > 0 {
        nodePar, nodeL, nodeR := q[0], q[1], q[2]
        q = q[3:]
        nodePar.Left = nodeR
        nodePar.Right = nodeL
        if nodeL != nil { 
        //if nodeL != nil && (node.Left != nil || node.Right != nil ) { 
            q = append(q, nodeL)
            q = append(q, nodeL.Left)
            q = append(q, nodeL.Right)
        }
        if nodeR != nil { 
        //if nodeR != nil && (node.Left != nil || node.Right != nil ) { 
            q = append(q, nodeR)
            q = append(q, nodeR.Left)
            q = append(q, nodeR.Right)
        }
        
    }
    return root
    
    
    
}