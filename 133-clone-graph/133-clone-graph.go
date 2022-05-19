


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */


func cloneGraph(node *Node) *Node {
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


/*
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
*/










    
//time : O(|V| + |E|)
//space: O(N) b/c of recursion stack and visited ad cloneNodeMap aux D.S.
/*
func cloneGraph(node *Node) *Node {
    if node == nil { return nil }
    
    //The number of nodes in the graph is in the range [0, 100].
    var visited [101]bool
    var cloneNodeMap [101]*Node  //cloneNodeMap := map[int]*Node{}
    
    var dfs func(node *Node)
    dfs = func(node *Node) {
        if visited[node.Val] { return }
        visited[node.Val] = true
        nodeClone := &Node{Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))}
        cloneNodeMap[node.Val] = nodeClone
        for _, srcNeighbor := range node.Neighbors {
            srcId := srcNeighbor.Val
            if !visited[srcId] {
                dfs(srcNeighbor)
            }
            neighborClone := cloneNodeMap[srcId] 
            nodeClone.Neighbors = append(nodeClone.Neighbors, neighborClone)
        }
    }
    
    if len(node.Neighbors) == 0 { 
        dfs(node)
        return cloneNodeMap[node.Val]
    }
    
    for _, srcNeighbor := range node.Neighbors {
        id := srcNeighbor.Val
        if !visited[id] {
            dfs(srcNeighbor)
        }
    }
    return cloneNodeMap[node.Val]
    
}
    */