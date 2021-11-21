//time: O(n) + O(nlogn) ~ O(nlogn)
//space: O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(a int) int {if a < 0 { return -a }; return a }
func min(a, b int) int { if a <= b {return a}; return b }

func minDiffInBST(root *TreeNode) int {
    maxInt32 := (1 << 31) -1
    minDiff := maxInt32
    lvlOrderVals := []int{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        lvlOrderVals = append(lvlOrderVals, node.Val)
        if node.Left != nil { q = append(q, node.Left)}
        if node.Right != nil { q = append(q, node.Right)}
    } 
    sort.Ints(lvlOrderVals)
    for i:=1; i<len(lvlOrderVals); i++ {
        diff := abs(lvlOrderVals[i] - lvlOrderVals[i-1])
        minDiff = min(minDiff, diff)
    }
    return minDiff
    
}