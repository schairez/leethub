/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS approach
// time: O(n)
// space: O(n)

func maxDepth(root *TreeNode) int {
    depth := 0
    if root == nil {
        return depth
    }
    var node *TreeNode
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        depth++
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return depth
}

















//postorder traversal
//time: O(n)
//space: O(n)
func maxDepthDFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    fromLeft := maxDepth(root.Left) 
    fromRight := maxDepth(root.Right) 
    if fromLeft >= fromRight {
        return fromLeft+1
    }
    return fromRight+1
}