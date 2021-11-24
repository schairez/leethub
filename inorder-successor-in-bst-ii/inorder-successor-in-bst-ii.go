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
    fmt.Println(node)
    in := node
    startNodeVal := in.Val
    var successor *Node
    Loop: 
    for node != nil {
        switch {
        case node.Right != nil:
            node = node.Right
            successor = node
            for successor.Left != nil {
                fmt.Println("here")
                successor = successor.Left
            }
            break Loop
        //case node.Val <= startNodeVal:
        case node.Parent != nil:
            for node.Val <= startNodeVal && node.Parent != nil {
                node = node.Parent
            }
            if node.Val > startNodeVal {
                successor = node
            }
            break Loop
        case node.Right == nil:
            break Loop
        case node.Val > startNodeVal:
            successor = node
            if node.Left != nil {
                successor = node.Left
            }
            break Loop
        }
    } 
    return successor
}