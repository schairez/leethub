/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//time: O(n)
//space: O(n)


//import "fmt"

func findLeaves(root *TreeNode) [][]int {
    res := make([][]int,1)
    var postorder func(root *TreeNode, idx int) int
    postorder = func(root *TreeNode, idx int) int {
        if root == nil { return 0 }
        
        idxL := postorder(root.Left, idx)
        idxR := postorder(root.Right, idx)
        idxNew := max(idxL, idxR)
        
        if len(res) == idxNew {
            res = append(res, []int{})
        } 
        //fmt.Println(res)
        res[idxNew] = append(res[idxNew], root.Val)
        root.Left = nil
        root.Right = nil
        return idxNew+1
    }
    
    postorder(root, 0)
    return res
}

func max(a, b int) int {
    if a >= b { return a}
    return b
}