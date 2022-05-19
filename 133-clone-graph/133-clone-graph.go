


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
    var (
        currNode *Node
        queue []*Node
        nodeMap map[*Node]*Node
    )
    nodeMap = make(map[*Node]*Node)
    srcNodeCopy := &Node{Val: node.Val}
    nodeMap[node] = srcNodeCopy
    
    queue = append(queue, node)
    for len(queue) != 0 {
        currNode, queue = queue[0], queue[1:]
        for _, nei := range currNode.Neighbors {
            if _, seen := nodeMap[nei]; !seen {
                nodeMap[nei] = &Node{Val: nei.Val}
                queue = append(queue, nei)
            }
            currNodeCopy := nodeMap[currNode]
            currNodeCopy.Neighbors = append(currNodeCopy.Neighbors, nodeMap[nei])
        }
    }
    return srcNodeCopy
}

    /*
func cloneGraph(node *Node) *Node {
    // Node.val is unique for each node.
    var (
        visited   map[int]*Node
        queue     [][2]*Node 
        currNodes [2]*Node
    )
    srcNodeCopy := &Node{}
    visited[node.Val] = srcNodeCopy
    queue = append(queue, [2]*Node{node, srcNodeCopy})
    for len(queue) != 0 {
        currNodes, queue = queue[0], queue[1:]
        origNode, nodeCopy := currNodes[0], currNodes[1]
        nodeCopy.Val = origNode.Val
        for _, nei := range origNode.Neighbors {
            if _, seen := visited[nei.Val]; !seen {
                visited[nei.Val] 
            }
            neiCopy := &Node{}
            nodeCopy.Neighbors = append(nodeCopy.Neighbors, neiCopy)
            queue = append(queue, [2]*Node{nei, neiCopy})
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