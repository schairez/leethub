/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    n := len(preorder)
    locMap := make(map[int]int, n)
    for idx, val := range inorder {
        locMap[val] = idx
    }
    var (
        currIdx int
        dfs     func(int, int) *TreeNode
    )
    
    dfs = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }
        
        val := preorder[currIdx]
        currIdx++
        root := &TreeNode{val, nil, nil}
        root.Left = dfs(left, locMap[val]-1)
        root.Right = dfs(locMap[val]+1, right)
        return root
    }
    return dfs(0, n-1)
}



