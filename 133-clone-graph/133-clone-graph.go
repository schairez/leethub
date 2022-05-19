/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    nodeMap := make(map[*Node]*Node)
    
    var dfs func(*Node) *Node
    dfs = func(orig *Node) *Node {
        if nodeCopy, exists := nodeMap[orig]; exists {
            return nodeCopy
        }
        nodeCopy := &Node{Val: orig.Val}
        nodeMap[orig] = nodeCopy
        for _, nei := range orig.Neighbors {
            nodeCopy.Neighbors = append(nodeCopy.Neighbors, dfs(nei))
        }
        return nodeCopy
    }
    
    return dfs(node)
}