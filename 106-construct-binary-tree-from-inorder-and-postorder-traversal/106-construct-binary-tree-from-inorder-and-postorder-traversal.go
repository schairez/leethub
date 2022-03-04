/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(inorder)
    currIdx := n-1
    locMap := make(map[int]int, n)
    for idx, val := range inorder {
        locMap[val] = idx
    }
    var dfs func(int, int) *TreeNode
    // reverse postorder
    dfs = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }
        val := postorder[currIdx]
        root := &TreeNode{val, nil, nil}
        currIdx--
        root.Right = dfs(locMap[val]+1, right)
        root.Left = dfs(left, locMap[val]-1)
        return root
    }
    
    return dfs(0, n-1)
    
}