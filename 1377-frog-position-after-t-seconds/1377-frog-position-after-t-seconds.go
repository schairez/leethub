// 1377. Frog Position After T Seconds
// time: O(V+E)
// space: O(V+E)

func frogPosition(n int, edges [][]int, t int, target int) float64 {
    graph := make([]map[int]struct{}, n+1) 
    for _, edge := range edges {
        src, dst := edge[0], edge[1]
        if graph[src] == nil {
            graph[src] = make(map[int]struct{})
        }
        graph[src][dst] = struct{}{}
        if graph[dst] == nil {
            graph[dst] = make(map[int]struct{})
        }
        graph[dst][src] =struct{}{}
    }
    type nodeData struct {
        nodeId int
        prob float64
    }
    secondsLeft := t
    queue := make([]nodeData, 0, n+1)
    // frog starts at vertex 1
    queue = append(queue, nodeData{1, 1.0})
    var data nodeData
    for len(queue) != 0 && secondsLeft >= 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            data, queue = queue[0], queue[1:]
            if data.nodeId == target && secondsLeft == 0 {
                return data.prob
            }
            numNodes := len(graph[data.nodeId])
            if numNodes > 0 {
                nextProb := (1.0 / float64(numNodes)) * data.prob
                for neiId := range graph[data.nodeId] {
                    queue = append(queue, nodeData{neiId, nextProb})
                    delete(graph[data.nodeId], neiId)
                    delete(graph[neiId], data.nodeId)
                }
            } else {
                if data.nodeId == target {
                    queue = append(queue, data)
                }
            }
        }
        secondsLeft--
    }
    return 0.0
}