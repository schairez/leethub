/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//rev inorder recursive approach
//RDL
//reversed inorder (RDL) 
/*
space: O(height); if skewed tree height = depth = numNodes ~ O(n)
if balanced tree space complexity os omega(height) = O(logn)
time: O(numNodes) -> O(n)
*/

func convertBST(root *TreeNode) *TreeNode {
    sumV := 0
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Right)
        sumV += node.Val
        node.Val = sumV
        dfs(node.Left)
    }
    refNode := root
    dfs(root)
    return refNode
    
}

// time: O(n); size of BST
// space: O(n) space b/c of stack

//reverse in-order iterative stack 

func convertBSTIterative(root *TreeNode) *TreeNode {
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