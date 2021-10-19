/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: O(n); size of BST
// space: O(n) space b/c of stack

//reverse in-order iterative stack 

func convertBST(root *TreeNode) *TreeNode {
    var stack []*TreeNode
    curr := root
    sum := 0 
    for {
        if curr != nil {
            stack = append(stack, curr)
            curr = curr.Right
        } else if len(stack) > 0 {
            curr = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            sum += curr.Val
            curr.Val = sum
            curr = curr.Left
        } else {
            break
        }
    }
    return root
}