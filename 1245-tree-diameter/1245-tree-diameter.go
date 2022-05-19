// 1245. Tree Diameter
// for the tree we have N-nodes and N-1 edges
// to build the tree's adjList we iter through O(N) edge pairs and form a space of O(n)
// we call BFS twice, our logic visits each node once and use a queue and hashmap, thus O(n) space and time for each
// of the two runs
// thus, time: O(n); space: O(n)

type nodeData struct {
    nodeId int
    dist   int
}

func treeDiameter(edges [][]int) int {
    n := len(edges)
    graph := make([][]int, n+1)
    for _, edge := range edges {
        src, dst := edge[0], edge[1]
        graph[src] = append(graph[src], dst)
        graph[dst] = append(graph[dst], src)
    }
    // locate a node at the last level from srcId=0
    nodeData1 := bfs(graph, 0)
    // locate peripheral node from last lvl node and calc distance from u to v 
    nodeData2 := bfs(graph, nodeData1.nodeId)
    diameter := nodeData2.dist
    
    return diameter
}



func bfs(graph [][]int, srcNode int) nodeData { 
    n := len(graph)
    var (
        queue []int
        node  int
    )
    lvl := 0
    visited := make(map[int]struct{}, n)
    visited[srcNode] = struct{}{}
    queue = append(queue, srcNode)
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            for _, neiNode := range graph[node] {
                if _, seen := visited[neiNode]; !seen {
                    visited[neiNode] = struct{}{}
                    queue = append(queue, neiNode)
                }
            }
        }
        if len(queue) != 0 {
            lvl++
        }
    }
    return nodeData{node, lvl}
}

