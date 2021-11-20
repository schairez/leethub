/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Two nodes of a binary tree are cousins 
//if they have the same depth with different parents
//Each node has a unique value.

func isCousins(root *TreeNode, x int, y int) bool {
    var foundX, foundY bool
    var xDepth, yDepth int 
    var xParent, yParent int
    
    postOrder(root, root, 0, func(nodeV, parentV, curDepth int) bool {
        switch nodeV {
        case x:
            foundX = true
            xDepth = curDepth
            xParent = parentV
        case y:
            foundY = true
            yDepth = curDepth
            yParent = parentV
        }
        if foundX == true && foundY == true {
            return true
        }
        return false
    })
    
    return xDepth == yDepth && xParent != yParent
}

//post -> LRD
func postOrder(node *TreeNode, parent *TreeNode, curDepth int, f func(node, parentV, curDepth int) bool) {
    if node != nil {
        postOrder(node.Left, node, curDepth+1, f)
        postOrder(node.Right, node, curDepth+1, f)
        foundBoth := f(node.Val, parent.Val, curDepth)
        if foundBoth { return }
    }
}