/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil { return root }
    queue := []*Node{}
    queue = append(queue, root)
    for len(queue) > 0 {
        currLen := len(queue)
        for i:=0; i < currLen; i++ {
            currNode := queue[0]
            queue = queue[1:]
            if i < currLen - 1 {
                currNode.Next = queue[0]
            }
            if currNode.Left != nil {
                queue = append(queue, currNode.Left)
            }
            if currNode.Right != nil {
                queue = append(queue, currNode.Right)
            }
        }
    }
    return root
}