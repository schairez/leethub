/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//morris traversal
//time: O(2n) ~ O(n)
//space: O(heightTree) ~ O(n)

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    var (
        currNode *TreeNode 
        predNode *TreeNode 
    )
    res := []int{}
    currNode = root
    for currNode != nil {
        if currNode.Left != nil {
            predNode = currNode.Left
            for predNode.Right != nil && predNode.Right != currNode {
                predNode = predNode.Right
            }
            if predNode.Right == nil {
                predNode.Right = currNode
                currNode = currNode.Left
            } else {
                res = append(res, currNode.Val)
                predNode.Right = nil
                currNode = currNode.Right
            }
        } else {
            res = append(res, currNode.Val)
            currNode = currNode.Right
        }
    }
    return res
}




//inorder -> LDR

//iterative
//time:O(n); space:O(n)

func inorderTraversalStackIter(root *TreeNode) []int {
    res := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, root.Val)
        root = root.Right
    }
    return res
}




//trivial recursive v
//space:O(n)
//time:O(n)
func inorderTraversalR(root *TreeNode) []int {
    if root == nil { return []int{} }
    res := []int{}
    var helper func(node *TreeNode) 
    helper = func(node *TreeNode) {
        if node == nil { return }
        //LDR; inorder
        helper(node.Left)
        res = append(res, node.Val)
        helper(node.Right)
    } 
    helper(root)
    
    return res
}
