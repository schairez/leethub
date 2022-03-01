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
    if root == nil {
        return root
    }
    queue := []*Node{root}
    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            if i > 0 {
                queue[i-1].Next = queue[i]
            }
            if lChild := queue[i].Left; lChild != nil {
                queue = append(queue, lChild)
            }
            if rChild := queue[i].Right; rChild != nil {
                queue = append(queue, rChild)
            }
        }
        queue = queue[n:]
    }
    
    return root
	
}