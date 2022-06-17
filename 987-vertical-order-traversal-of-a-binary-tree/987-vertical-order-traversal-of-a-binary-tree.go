/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 987. Vertical Order Traversal of a Binary Tree
// time: O(nlogn) to sort
// space: O(n)

func verticalTraversal(root *TreeNode) [][]int {
    type pair struct {val, row int}
    type data struct { node *TreeNode; row, col int }
    // m[col] = [{val, rowI}, {val, rowI},...]
    nodesMap := make(map[int][]pair)
    minCol := 0
    var (
        queue []data
        currData data
    )
    queue = append(queue, data{root, 0, 0})
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            currData, queue = queue[0], queue[1:]
            node, row, col := currData.node, currData.row, currData.col
            nodesMap[col] = append(nodesMap[col], pair{node.Val, row})
            if node.Left != nil {
                if minColCand := col-1; minColCand < minCol {
                    minCol = minColCand
                }
                queue = append(queue, data{node.Left, row+1, col-1 })
            }
            if node.Right != nil {
                queue = append(queue, data{node.Right, row+1, col+1})
            }
        }
    }
    n := len(nodesMap)
    res := make([][]int, 0, n)
    currLen, i := 0, minCol
    for currLen < n {
        colData := nodesMap[i]
        sort.Slice(colData, func(a, b int) bool {
            data1, data2 := colData[a], colData[b]
            if data1.row == data2.row {
                return data1.val < data2.val
            }
            return data1.row < data2.row
        })
        tmp := make([]int, 0, len(colData))
        for _, elem := range colData {
            tmp = append(tmp, elem.val)
        }
        res = append(res, tmp)
        currLen++
        i++
    }
    
    return res
}

