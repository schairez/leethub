/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

// time: O(n)
// space: O(n)
func treeToDoublyList(root *Node) *Node {
    if root == nil {
        return root
    }
    var head, tail *Node
    var dfs func(*Node)
    dfs = func(node *Node) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if head == nil {
            head = node
        }
        if tail != nil {
            node.Left = tail
            tail.Right = node
        }
        tail = node
        dfs(node.Right)
    }
    dfs(root)
    head.Left = tail
    tail.Right = head
    return head
}