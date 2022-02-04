/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//iterative approach
//morris traversal approach
//space optimized version
//overkill for this problem x_X
// time: O(n)
//space: O(1)

func kthSmallest(root *TreeNode, k int) int {
    kthVal := k
    var (
        res int
    )
    checkKthFn := func(val int) {
        kthVal--
        if kthVal == 0 {
            res = val
        }
    }
    morrisTraversal(root, checkKthFn)
    return res
}


func morrisTraversal(root *TreeNode, checkKthFn func(val int)) {
    var (
        successorNode *TreeNode
        predNode *TreeNode
    )
    successorNode = root
    for successorNode != nil {
        if successorNode.Left != nil {
            predNode = successorNode.Left
            for predNode.Right != nil && predNode.Right != successorNode {
                predNode = predNode.Right
            }
            if predNode.Right == nil {
                predNode.Right = successorNode
                successorNode = successorNode.Left
            } else {
                checkKthFn(successorNode.Val)
                predNode.Right = nil
                successorNode = successorNode.Right
            }
        } else {
            checkKthFn(successorNode.Val)
            successorNode = successorNode.Right
        }
    }
}








//time: O(logn + k) if balanced tree; O(n+k) if skewed 
//space: O(logn + k) if balanced tree; O(n+k) if skewed 

//recursive
func kthSmallestRec(root *TreeNode, k int) int {
    kthVal := k
    var res int
    inorder(root, func(val int) bool {
        kthVal--
        if kthVal == 0 {
            res = val
            return true
        }
        return false
    })
    
    return res
}

func inorder(node *TreeNode, checkFn func(val int) bool) {
    if node == nil {
        return
    }
    inorder(node.Left, checkFn)
    if done := checkFn(node.Val); done == true {
        return
    }
    inorder(node.Right, checkFn)
}