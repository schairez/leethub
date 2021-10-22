/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time: O(NlogN)
// space: O(1)

func bstFromPreorder(preorder []int) *TreeNode {
    //var curr *TreeNode
    root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
    //curr = root
    for i:=1; i < len(preorder); i++ {
        arrVal := preorder[i]
        curr := root
        newNode :=  &TreeNode{Val: arrVal, Left: nil, Right: nil} 
        for curr != nil { //time: log(N) 
            if arrVal < curr.Val {
                if curr.Left == nil { 
                    curr.Left = newNode
                    break
                } 
                curr = curr.Left
            } else {
                if curr.Right == nil {
                    curr.Right = newNode
                    break
                }
                curr = curr.Right
            }
        }
    }
    
    return root
}