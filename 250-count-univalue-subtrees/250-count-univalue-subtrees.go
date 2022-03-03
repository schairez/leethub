/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) int {
    // dfs postorder
    var dfs func(node *TreeNode) (int, bool)
    dfs = func(node *TreeNode) (int, bool) {
        if node == nil {
            return 0, true
        }
        uniLeftCnt, isUniLeft := dfs(node.Left)
        uniRightCnt, isUniRight := dfs(node.Right) 
        total := uniLeftCnt + uniRightCnt
        if isUniLeft && isUniRight {
            if node.Left != nil && node.Left.Val != node.Val {
                return total, false
            } 
            if node.Right != nil && node.Right.Val != node.Val {
                return total, false
            }
            return 1 + total, true
        }
        return total, false
    }
    
    res, _ := dfs(root)
    return res
    
}