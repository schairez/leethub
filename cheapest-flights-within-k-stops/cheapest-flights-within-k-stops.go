import (
    "math"
    "container/heap"
    "fmt"
)

/*
number of levels to be explored is bounded by K
new_dist = dist[u] + length(u, v) 
if new_dist < dist[v]:
    dist[v] = new_dist

*/

type Edge struct {
    DstId int
    Weight int
}



func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    min := func(a,b int) int {
        if a <= b { return a}
        return b
    }
    graph := make(map[int][]Edge, n)
    stops := make([]int, n) //minimize number of stops, s.t. K
    dist := make([]int, n) //minimize distance dV = min(dU + wUV, dV)
    for i:=0; i < n; i++ {
        dist[i] = int(math.MaxInt64)
        stops[i] = int(math.MaxInt64)
    }
    
    for i := range flights {
        u := flights[i][0]
        if _, ok := graph[u]; !ok {
            graph[u] = make([]Edge, 0)
        }
        v, wUV := flights[i][1], flights[i][2]
        graph[u] = append(graph[u], Edge{DstId: v, Weight: wUV})
    }
    
    dist[src] = 0
    stops[src] = 0
    
    fmt.Println(graph)
    //pq := &PQ{}
    pq := make(PQ, 0, n)
    heap.Init(&pq)
    heap.Push(&pq, Node{Id: src, PriceSoFar: 0, Stops: 0})
    for pq.Len() > 0 {
        fmt.Println(pq)
        city := heap.Pop(&pq).(Node)
        fmt.Println(city.Id)
        if city.Id == dst {
            return city.PriceSoFar
        }
        if city.Stops == k + 1 {
            continue
        }
        dU := dist[city.Id]
        numStops := city.Stops
        neighborEdges := graph[city.Id]
        for _, edge := range neighborEdges {
            dV, wUV := dist[edge.DstId], edge.Weight
            addToHeap := dU+wUV < dV || numStops < stops[edge.DstId] 
            //if (dU + wUV < dV) //minimize cost distance
            //if stops < stops[dstId] // minimize number of stops
            dist[edge.DstId] = min(dist[edge.DstId], dU+wUV)
            if addToHeap {
                heap.Push(&pq, Node{
                        Id: edge.DstId,
                        PriceSoFar: city.PriceSoFar+ edge.Weight,
                        Stops: city.Stops +1,
                    })
            }
            stops[edge.DstId] = numStops 
        }
    }
    return -1
}

type Node struct {
    Id int
    PriceSoFar int
    Stops int
}

type PQ []Node

func (pq PQ) Len() int { return len(pq)}

func (pq PQ) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less (i, j int) bool {
    return pq[i].PriceSoFar < pq[j].PriceSoFar
}

func (pq *PQ) Push(x interface{}) {
    *pq = append(*pq, x.(Node))
}


func (pq *PQ) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[:n-1]
    return x
}



