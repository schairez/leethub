func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    // return all MHT's root labels
    inDegree := make([]int, n)
    graph := make([][]int, n)
    for _, edge := range edges {
        src, dst := edge[0], edge[1]
        graph[src] = append(graph[src], dst)
        graph[dst] = append(graph[dst], src)
        inDegree[src]++
        inDegree[dst]++
    }
    var (
        leafQueue  []int
        node       int
    )
    for i := range inDegree {
        if inDegree[i] == 1 {
            leafQueue = append(leafQueue, i)
        }
    }
    var res []int
    for len(leafQueue) != 0 { 
        res = []int{}
        for currLen := len(leafQueue); currLen != 0; currLen-- {
            node, leafQueue = leafQueue[0], leafQueue[1:]
            res = append(res, node)
            for _, nei := range graph[node] {
                inDegree[nei]--
                if inDegree[nei] == 1 {
                    leafQueue = append(leafQueue, nei)
                }
            }
        }
    }
    
    return res
}