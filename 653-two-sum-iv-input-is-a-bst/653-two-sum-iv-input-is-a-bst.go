//time:  O(n)
//space: O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//inorder
//LDR
//recursive call stack O(n) space O(n) time
func inOrderDFS(root *TreeNode, f func(val int)) {
    if root == nil {
        return
    }
    inOrderDFS(root.Left, f)
    f(root.Val)
    inOrderDFS(root.Right, f)
}

func findTarget(root *TreeNode, k int) bool {
    var inOrder []int 
    inOrderDFS(root, func(val int) {
        inOrder = append(inOrder, val)
    })
    n := len(inOrder)
    left, right := 0, n-1
    for left != right {
        sumVals := inOrder[left] + inOrder[right]
        if sumVals == k {
            return true
        } else if sumVals > k {
            right--
        } else {
            left++
        }
    }
    return false
}


