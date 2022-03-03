/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time: O(n)
// space: O(1)
// overkill morris traversal ;)


func isUnivalTree(root *TreeNode) bool {
    var (
        currNode *TreeNode
        predNode *TreeNode //predecessor
    )
    isUnival := true 
    currNode = root
    uniVal := currNode.Val 
    for currNode != nil {
        if currNode.Val != uniVal {
            isUnival = false
        }
        if currNode.Left != nil {
            predNode = currNode.Left
            for predNode.Right != nil && predNode.Right != currNode {
                predNode = predNode.Right
            }
            if predNode.Right == nil {
                predNode.Right = currNode
                currNode = currNode.Left
            } else {
                predNode.Right = nil
                currNode = currNode.Right
            }
        } else {
            currNode = currNode.Right
        }
    }
    
    return isUnival
}