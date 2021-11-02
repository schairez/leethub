/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    var res []int
    
    if root == nil { return res }
    
    var curr *TreeNode
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        curr = stack[len(stack)-1]
        res = append(res, curr.Val)
        stack = stack[:len(stack)-1]
        if curr.Left != nil { stack = append(stack, curr.Left) }
        if curr.Right != nil { stack = append(stack, curr.Right) }

    }
    revInPlace(&res)    
    
    return res

}
func revInPlace(nums *[]int) {
    for i, j := 0,len(*nums)-1; i <j; i,j=i+1,j-1 {
        (*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i] 
    }
}




//postorder -> LRD
//time:O(n); space: O(n)
//naive recursive v
func postorderTraversalR(root *TreeNode) []int {
    if root == nil { return []int{} }
    res := []int{}
    var helper func(node *TreeNode)
    helper = func(node *TreeNode) {
        if node == nil { return }
        //LRD
        helper(node.Left)
        helper(node.Right)
        res = append(res, node.Val)
    }
    helper(root)
    return res
}