/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// time: O(n)
// space: O(1)
func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    lvlStartNode := root
    for lvlStartNode != nil {
        var prevNode *Node
        currNode := lvlStartNode
        lvlStartNode = nil
        for currNode != nil {
            if currNode.Left != nil {
                if prevNode != nil {
                    prevNode.Next = currNode.Left
                } else {
                    lvlStartNode = currNode.Left
                }
                prevNode = currNode.Left
            }
            if currNode.Right != nil {
                if prevNode != nil {
                    prevNode.Next = currNode.Right
                } else {
                    lvlStartNode = currNode.Right
                }
                prevNode = currNode.Right
            }
            currNode = currNode.Next
        }
    }
    return root
}

// BFS approach
// time: O(n)
// space: O(maxWidthLvl) ≈ O(n)

func connectBFS(root *Node) *Node {
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