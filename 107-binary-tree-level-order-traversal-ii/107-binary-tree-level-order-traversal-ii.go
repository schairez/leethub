/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//time: O(numNodes) + O(rev) ≈ O(n)
//space: O(numNodes at maxLvl) ≈ O(m)

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    var res [][]int
    queue := []*TreeNode{root} 
    var node *TreeNode
    for len(queue) != 0 {
        n := len(queue)
        nodeVals := make([]int, 0, n)
        for i := 0; i < n; i++ {
            node, queue = queue[0], queue[1:]
            nodeVals = append(nodeVals, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, nodeVals)
    }
    
    rev2DSlice(res)
    
    return res
}

func rev2DSlice(valsAtLvl [][]int) {
    n := len(valsAtLvl)
    start, end := 0, n-1
    for start < end {
        valsAtLvl[start], valsAtLvl[end] = valsAtLvl[end], valsAtLvl[start]
        start++
        end--
    }
}