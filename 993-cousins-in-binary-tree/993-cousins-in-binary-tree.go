

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS approach
func isCousins(root *TreeNode, x int, y int) bool {
    type nodeWParent struct {
        *TreeNode
        parent int
    }
    var (
        queue []nodeWParent
        node  nodeWParent
    )
    queue = append(queue, nodeWParent{root, -1})
    for len(queue) != 0 {
        parent := make([]int, 0, 2)
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node.Val == x || node.Val == y {
                parent = append(parent, node.parent)
            }
            if len(parent) == 2 {
                return parent[0] != parent[1]
            }
            if node.Left != nil {
                queue = append(queue, nodeWParent{node.Left, node.Val})
            }
            if node.Right != nil {
                queue = append(queue, nodeWParent{node.Right, node.Val})
            }
        }
    }
    return false
}












// DFS approach
//Two nodes of a binary tree are cousins 
//if they have the same depth with different parents
//Each node has a unique value.
//time: O(n)
//space: O(n)

func isCousinsDFS(root *TreeNode, x int, y int) bool {
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