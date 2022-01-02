/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//time: O(n) worst case, best case O(logn) if balanced BST
//space: O(1)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        newNode := &TreeNode{Val: val}
        return newNode
    }
    currNode := root
    for currNode.Val != val {
        if currNode.Val < val {
            rightNextNode := currNode.Right
            if rightNextNode == nil {
                currNode.Right = &TreeNode{Val: val}
            }
            currNode = currNode.Right
        } else {
            leftNextNode := currNode.Left
            if leftNextNode == nil {
                currNode.Left = &TreeNode{Val: val}
            }
            currNode = currNode.Left
        }
    }
    return root
}
//if balanced tree height is logn = H
//recursive call stack time and space avg case: O(log(n))
//worst case: O(n)
func insertIntoBSTRecursive(root *TreeNode, val int) *TreeNode {
    if root == nil {
        newNode := &TreeNode{Val: val}
        return newNode
    }
    if root.Val < val {
        root.Right = insertIntoBSTRecursive(root.Right, val)
    } else {
        root.Left = insertIntoBSTRecursive(root.Left, val)
    }
    return root
}
