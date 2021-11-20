/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//time: O(n)
//space: O(n)

type NodeData struct {
    Node *TreeNode
    Idx int
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := [][]int{}
    
    q := []NodeData{NodeData{root, 0}}
    for len(q) > 0 {
        nodeData := q[0]
        q = q[1:]
        val, idx := nodeData.Node.Val, nodeData.Idx
        if len(res) == idx {
            res = append(res, []int{})
        }
        res[idx] = append(res[idx], val)
        
        if nodeData.Node.Left != nil {
            q = append(q, NodeData{nodeData.Node.Left, idx+1})
        }
        if nodeData.Node.Right != nil {
            q = append(q, NodeData{nodeData.Node.Right, idx+1})
        }
    } 
   return res 
}