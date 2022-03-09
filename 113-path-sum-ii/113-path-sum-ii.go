/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 113. Path Sum II
// time: O(n)
// space: O(n)

func pathSum(root *TreeNode, targetSum int) [][]int {
    var (
        res   [][]int
        dfs   func(*TreeNode, []int, int)
    )
    
    dfs = func(node *TreeNode, pathSoFar []int, sumSoFar int) {
        if node == nil {
            return
        }
        sumSoFar += node.Val
        
        if node.Left == nil && node.Right == nil {
            pathSoFar = append(pathSoFar, node.Val)
            if sumSoFar == targetSum {
                tmp := make([]int, len(pathSoFar))
                copy(tmp, pathSoFar)
                res = append(res, tmp)
            }
            return
        }
        dfs(node.Left, append(pathSoFar, node.Val), sumSoFar)
        dfs(node.Right, append(pathSoFar, node.Val), sumSoFar)
    }
    
    dfs(root, []int{}, 0)
    
    return res
}