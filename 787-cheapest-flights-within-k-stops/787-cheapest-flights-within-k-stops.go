
//import "container/heap"

type Node struct {
    dstId, cost, numStops, idx int  
}

type PriorityQueue struct {
    data []*Node
    lessFn func(i, j int) bool
}

func newPQ(lessFn func(i, j int) bool) *PriorityQueue {
    return &PriorityQueue{
        lessFn: lessFn,
    }
}

func (pq *PriorityQueue) Len() int {
    return len(pq.data)
}
func (pq *PriorityQueue) Less(i , j int) bool {
    return pq.lessFn(pq.data[i].cost, pq.data[j].cost)
}
func (pq *PriorityQueue) Swap(i, j int) {
    pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
    pq.data[i].idx = i 
    pq.data[j].idx = j 
}
func (pq *PriorityQueue) Push(v interface{}) {
    node := v.(*Node)
    node.idx = pq.Len()
    pq.data = append(pq.data, node)
}
func (pq *PriorityQueue) Pop() interface{} {
    n := pq.Len()
    v := pq.data[n-1]
    pq.data[n-1] = nil
    v.idx = -1
    pq.data = pq.data[:n-1]
    return v
}
func (pq *PriorityQueue) Update(node *Node) {
    heap.Fix(pq, node.idx)
}

func min(a, b int) int {if a <= b {return a}; return b}

// prioritize cheapest cost from src to dst s.t. numStops <= k 

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    type edge struct { dstId, cost int }
    graph := make(map[int][]edge, n)
    // check to see if any edge links to dst node
    isDstLinkable := false
    maxInt32 := 1 << 31 - 1
    for _, flight := range flights {
        // flights[i] = [fromi, toi, pricei] 
        srcId, dstId, cost := flight[0], flight[1], flight[2]
        if _, exists := graph[srcId]; !exists {
            graph[srcId] = make([]edge, 0)
        }
        graph[srcId] = append(graph[srcId], edge{dstId, cost}) 
        if dstId == dst {
            isDstLinkable = true
        }
    }
    if !isDstLinkable {
        return -1
    }
    
    nodes := make([]*Node, n)
    for i := range nodes {
        nodes[i] = &Node{
            dstId: i,
            cost: maxInt32,
            numStops: maxInt32,
            idx: -1,
        }
    }
    pq := newPQ(func(a, b int) bool {
        return a < b
    })
    heap.Init(pq)
    srcNodeData := nodes[src]
    srcNodeData.cost, srcNodeData.numStops = 0, 0
    heap.Push(pq, srcNodeData)
    for pq.Len() != 0 {
        fmt.Println()
        nodeData := heap.Pop(pq).(*Node)
        nodeId, nodeCost := nodeData.dstId, nodeData.cost
        if nodeId == dst {
            return nodeData.cost
        }
        numStops := nodeData.numStops
        candNumStops := numStops + 1
        neiEdges := graph[nodeId]
        for _, neiEdge := range neiEdges {
            edgeCost, dstId := neiEdge.cost, neiEdge.dstId
            candCost := nodeCost + edgeCost 
            dstNodeData := nodes[dstId]
            isFinalDst := dstId == dst
            if !isFinalDst && candNumStops == k+1 {
                continue
            }
            if dstNodeData.idx == -1 || candCost < dstNodeData.cost || candNumStops< k+1 {
                if dstNodeData.idx == -1 {
                    dstNodeData.cost = candCost
                    dstNodeData.numStops = candNumStops
                    heap.Push(pq, dstNodeData)
                } else if candCost < dstNodeData.cost && candNumStops < dstNodeData.numStops { 
                    dstNodeData.cost = candCost
                    dstNodeData.numStops = candNumStops
                    pq.Update(dstNodeData)
                } else { // other cand node 
                    dstNodeDataCand := &Node{dstNodeData.dstId, candCost, candNumStops, -1}
                    heap.Push(pq, dstNodeDataCand)
                }
            }
        }
    }
    
    return -1
}



