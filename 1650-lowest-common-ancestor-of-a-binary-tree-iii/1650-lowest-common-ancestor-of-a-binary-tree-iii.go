/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

// time: O(n)
// space: O(n)

func lowestCommonAncestor(p *Node, q *Node) *Node {
    nodeSet := make(map[int]struct{})
    for node := p; node != nil; node = node.Parent {
        nodeSet[node.Val] = struct{}{}
    }
    for node := q; node != nil; node = node.Parent {
        if _, exists := nodeSet[node.Val]; exists {
            return node
        }
        nodeSet[node.Val] = struct{}{}
    }
    
    return nil
}
