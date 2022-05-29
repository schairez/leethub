/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func inorder(root *TreeNode) []*TreeNode {
    var res []*TreeNode
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        res = append(res, node)
        dfs(node.Right)
    }
    dfs(root)
    return res
}
func arrToBST(nodes []*TreeNode) *TreeNode {
    if len(nodes) == 0 {
        return nil
    }
    mid := len(nodes) >> 1
    retNode := nodes[mid]
    retNode.Left = arrToBST(nodes[:mid])
    retNode.Right = arrToBST(nodes[mid+1:])
    return retNode
}
    

func balanceBST(root *TreeNode) *TreeNode {
    nodes := inorder(root)
    return arrToBST(nodes)
}