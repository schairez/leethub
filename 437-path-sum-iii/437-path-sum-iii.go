/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 437. Path Sum III
// time: O(n)
// space: O(h) = O(logn) if balanced; else O(n)


func pathSum(root *TreeNode, targetSum int) int {
    valsMap := make(map[int]int)
    var (
        totalPaths int
        dfs        func(*TreeNode, int)
    )
    // postorder approach
    dfs = func(node *TreeNode, sumSoFar int) {
        if node == nil {
            return
        }
        sumSoFar += node.Val
        if sumSoFar == targetSum {
            totalPaths++
        }
        complement := sumSoFar - targetSum
        if tot, ok := valsMap[complement]; ok {
            totalPaths += tot
        }
        valsMap[sumSoFar]++
        dfs(node.Left, sumSoFar)
        dfs(node.Right, sumSoFar)
        valsMap[sumSoFar]--
    }
    
    dfs(root, 0)
    return totalPaths
    
}


