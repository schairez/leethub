/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// incr

type status = int

// enum
// c style
const (
    hasCamera status =  iota // coveredNoCam and cam here
    coveredNoCam      // coveredNoCam no cam
    coverMePlz // not coveredNoCam (plz help) 
)


func minCameraCover(root *TreeNode) int {
    var (
        res int
        dfs func(*TreeNode) status
    )
    dfs = func(node *TreeNode) status {
        if node == nil { //if null then its coveredNoCam
            return coveredNoCam
        }
        fromL := dfs(node.Left)
        fromR := dfs(node.Right)
        if fromL == coverMePlz || fromR == coverMePlz {
            res++
            return hasCamera
        }
        if fromL == hasCamera || fromR == hasCamera {
            return coveredNoCam
        }  
        return coverMePlz
    }
    
    if dfs(root) == coverMePlz {
        res++
    }
    return res
}