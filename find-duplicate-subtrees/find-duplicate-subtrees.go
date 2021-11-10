//time: O(n)
//space: O(n)

import (
    "fmt"
    "strconv"
    "strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    res := []*TreeNode{}
    pathMap := map[string]int{}
    //postorder
    var dfs func(node *TreeNode ) string
    dfs = func(node *TreeNode) string {
        if node == nil { return "#" }
        
        left, right := dfs(node.Left), dfs(node.Right)
        path := strings.Builder{}
        path.WriteString("L")
        path.WriteString(left)
        path.WriteString(strconv.Itoa(node.Val))
        path.WriteString("R")
        path.WriteString(right)
        
        pathStr := path.String()
        if _, ok := pathMap[pathStr]; !ok {
            pathMap[pathStr] = 1
        } else {
            pathMap[pathStr]++
        } 
        if pathMap[pathStr] == 2 {
            res = append(res, node)
        }
        return pathStr
    }
    dfs(root)
    fmt.Println(pathMap)
    return res
}

/*

//time: O(n)
//space: O(n)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    res := []*TreeNode{}
    pathMap := map[string]int{}
    //postorder
    var dfs func(node *TreeNode ) string
    dfs = func(node *TreeNode) string {
        if node == nil { return "#" }
        left := dfs(node.Left)
        right := dfs(node.Right)
        pathStr := "L" + left + strconv.Itoa(node.Val) + "R" + right
        //pathStr := left + string(node.Val) + right
        if _, ok := pathMap[pathStr]; !ok {
            pathMap[pathStr] = 1
        } else {
            pathMap[pathStr]++
        } 
        if pathMap[pathStr] == 2 {
            res = append(res, node)
        }
        return pathStr
    }
    dfs(root)
    fmt.Println(pathMap)
    return res
}

*/