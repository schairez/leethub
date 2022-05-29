// 1382. Balance a Binary Search Tree
// time: O(n)
// space: O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

    

func balanceBST(root *TreeNode) *TreeNode {
    nodes := inorder(root)
    return arrToBST(nodes, 0, len(nodes)-1)
}

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

func arrToBST(nodes []int, l, r int) *TreeNode {
    if l > r {
        return nil
    }
    mid := l + (r-l) >> 1
    return &TreeNode{
        Val: nodes[mid],
        Left: arrToBST(nodes, l, mid-1),
        Right: arrToBST(nodes, mid+1, r),
    }
}
