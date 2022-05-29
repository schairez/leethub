/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func inorder(root *TreeNode) []int {
    var res []int
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        res = append(res, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return res
}
func arrToBST(nodes []int) *TreeNode {
    if len(nodes) == 0 {
        return nil
    }
    mid := len(nodes) >> 1
    return &TreeNode{
        Val: nodes[mid],
        Left: arrToBST(nodes[:mid]),
        Right: arrToBST(nodes[mid+1:]),
    }
}
    

func balanceBST(root *TreeNode) *TreeNode {
    nodes := inorder(root)
    return arrToBST(nodes)
}