/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    if root == nil {
        return []int{}
    }
    // constraint: all the values Node.val are unique
    var res []int
    distMap := make(map[int]int)
    dfs(root, target.Val, distMap)
    findNodes(&res, distMap, root, k, 0)
    
    return res
}

func dfs(root *TreeNode, targetVal int, distMap map[int]int) {
    if root == nil {
        return
    }
    if root.Val == targetVal {
        distMap[root.Val] = 0
        return
    }
    dfs(root.Left, targetVal, distMap)
    if root.Left != nil {
        if dist, exists := distMap[root.Left.Val]; exists {
            distMap[root.Val] = dist + 1
            return
        }
    }
    dfs(root.Right, targetVal, distMap)
    if root.Right != nil {
        if dist, exists := distMap[root.Right.Val]; exists {
            distMap[root.Val] = dist + 1
            return
        }
    }
}


func findNodes(res *[]int, distMap map[int]int, root *TreeNode, k int, currDist int) {
    if root == nil {
    //if root == nil || currDist > k {
        return
    }
    if dist, exists := distMap[root.Val]; exists {
        currDist = dist
    } 
    if currDist == k {
        *res = append(*res, root.Val)
    }
    findNodes(res, distMap, root.Left, k, currDist+1)
    findNodes(res, distMap, root.Right, k, currDist+1)
}
