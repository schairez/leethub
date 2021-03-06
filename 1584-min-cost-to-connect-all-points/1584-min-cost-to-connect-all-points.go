
// 
// we first iterate through n points, which the index represents the vertex node adding each node to our nodes slice
// this takes O(V) time to iterate through input and O(logV) to heapify this point into our heap
// in our while condition, we pop from our heap and iterate through all candidate edges that are unvisited, checking whether
// this is a better cost than our initialized +inf to this node
// this is done in |E| time, to pop from our heap in logV time 
// time: O( |V|log|V| + |E|log|V|) ≈ O(|E|log|V|)
//  
// space: O(|V|) for heap memory and an additional O(|V|) for the nodes slice

const (
    inf = 1 << 31 - 1
)

func minCostConnectPoints(points [][]int) int {
    n := len(points)
    nodes := make([]*PrimNode, n)
    pq := make(PriorityQueue, 0 )
    heap.Init(&pq)
    for idx := range points {
        node := &PrimNode{
            vertexId: idx,
            cost: inf,
            visited: false,
        }
        heap.Push(&pq, node)
        nodes[idx] = node
    }
    nodes[0].cost = 0
    heap.Fix(&pq, nodes[0].index) //O(logn) operation
    
    numVisited := 0
    minCost := 0
    for numVisited < n && pq.Len() != 0 {
        v1 := heap.Pop(&pq).(*PrimNode)
        minCost += v1.cost
        numVisited++
        v1.visited = true
        srcX, srcY := points[v1.vertexId][0], points[v1.vertexId][1]
        for nei := 0; nei < n; nei++ {
            v2 := nodes[nei]
            if v2.visited {
                continue
            }
            dstX, dstY := points[v2.vertexId][0], points[v2.vertexId][1]
            //cost of connecting two points
            if cost := manhattanDist(srcX, srcY, dstX, dstY); cost < nodes[v2.vertexId].cost {
                v2.cost = cost
                v2.predId = v1.vertexId
                heap.Fix(&pq, v2.index)
            }
        }
    }
    return minCost
}



func abs(a int) int { if a < 0 { return -a}; return a }

func manhattanDist(srcX, srcY, dstX, dstY int) int {
    return abs(srcX - dstX) + abs(srcY - dstY)
}

type PrimNode struct {
    vertexId         int     // a reference to the idx in our points input
    cost             int        // dist we cache and minimize to connect to this coordinate; cost to this node
    index            int        // idx of loc on the heap
    visited          bool
    predId           int        //predecessor id
}

type PriorityQueue []*PrimNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    node := x.(*PrimNode)
    node.index = len(*pq)
    *pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    node := (*pq)[n-1]
    (*pq)[n-1] = nil
    *pq = (*pq)[0 : n-1]  
    return node
}


