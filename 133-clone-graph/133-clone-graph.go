/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// 133. Clone Graph
//time : O(|V| + |E|)
// space: O(|V|)
// bfs approach
func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }
    // Node.val is unique for each node.
    nodeMap := make(map[int]*Node)
    srcNodeCopy := &Node{Val: node.Val}
    nodeMap[node.Val] = srcNodeCopy
    var (
        origNode *Node
        queue []*Node
    )
    queue = append(queue, node)
    for len(queue) != 0 {
        origNode, queue = queue[0], queue[1:]
        for _, nei := range origNode.Neighbors {
            if _, seen := nodeMap[nei.Val]; !seen {
                nodeMap[nei.Val] = &Node{Val: nei.Val}
                queue = append(queue, nei)
            }
            currNodeCopy := nodeMap[origNode.Val]
            currNodeCopy.Neighbors = append(currNodeCopy.Neighbors, nodeMap[nei.Val])
        }
    }
    return srcNodeCopy
}

    
//time : O(|V| + |E|)
//space: O(H) ≈ O(N) b/c of recursion stack and visited ad cloneNodeMap aux D.S.
// DFS approach
func cloneGraphDFS(node *Node) *Node {
    if node == nil {
        return node
    }
    nodeMap := make(map[*Node]*Node)
    
    var dfs func(*Node) *Node
    dfs = func(origNode *Node) *Node {
        if origNode == nil {
            return nil
        }
        if nodeCopy, exists := nodeMap[origNode]; exists {
            return nodeCopy
        }
        nodeCopy := &Node{ Val: origNode.Val}
        nodeMap[origNode] = nodeCopy
        for _, nei := range origNode.Neighbors {
            nodeCopy.Neighbors = append(nodeCopy.Neighbors, dfs(nei))
        }
        return nodeCopy
    }
    
    return dfs(node)
}