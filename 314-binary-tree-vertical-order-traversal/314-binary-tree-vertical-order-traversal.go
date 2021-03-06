/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func max(a, b int) int { if a >= b { return a}; return b}
func min(a, b int) int { if a <= b { return a}; return b}


//k = numKeys in map; v = num vertices per col key
//time: O(k * vlog(v))
//space: O(n) where n = numNodes; dfs stack calls 

type data struct {
    row int
    val int
}

func verticalOrder(root *TreeNode) [][]int {
    valsMap := make(map[int][]data)
    minKey, maxKey := 1 <<31 -1, -1 * 1<<31
    
    //dfs space: O(height) ~ O(n); time: O(n)
    dfs(valsMap, root, 0, 0, func(col int) {
        minKey = min(minKey, col)
        maxKey = max(maxKey, col)
    })
    res := make([][]int, len(valsMap))
    
    idx := 0
    //time: O(k) to iter keys of map * time to sort vals at col ~ O(k * colVals(log(colVals)))
    //therefore, time: O(k * vlog(v) + k * len(v)) ~ O(k * vlog(v))
    //space: O(logv) to sort
    for k := minKey; k <= maxKey; k++ {
        sort.SliceStable(valsMap[k], func(i, j int) bool {
            if valsMap[k][i].row < valsMap[k][j].row {
                return true
            }
            return false
        })
        for _, nodeData := range valsMap[k] {
            res[idx] = append(res[idx], nodeData.val)
        }
        idx++
    }
    
    return res
}

func dfs(valsMap map[int][]data, node *TreeNode, col, row int, updateKeysFn func(col int)) {
    if node == nil {
        return
    }
    if _, ok := valsMap[col]; !ok {
        valsMap[col] = []data{}
    }
    valsMap[col] = append(valsMap[col], data{row, node.Val})
    updateKeysFn(col)
    
    dfs(valsMap, node.Left, col-1, row+1, updateKeysFn)
    dfs(valsMap, node.Right, col+1, row+1, updateKeysFn)
    
}


/*
from left to right
smallest col is 1st in res 

postorder DFS (LRD)
each subsequent child call we +1 to row.

*/