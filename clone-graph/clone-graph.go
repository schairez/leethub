/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */


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
    
