/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func inorderSuccessor(node *Node) *Node {
    startNodeVal := node.Val
    var successor *Node
    switch {
    case node.Right != nil:
        node = node.Right
        successor = node
        for successor.Left != nil {
            successor = successor.Left
        }
    case node.Parent != nil:
        for node.Val <= startNodeVal && node.Parent != nil {
            node = node.Parent
        }
        if node.Val > startNodeVal {
            successor = node
        }
    default: //case node.Right == nil
        break
    } 
    return successor
}