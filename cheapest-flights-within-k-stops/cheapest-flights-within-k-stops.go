import (
    "math"
    "container/heap"
)

/*
number of levels to be explored is bounded by K
new_dist = dist[u] + length(u, v) 
if new_dist < dist[v]:
    dist[v] = new_dist

*/

type Node struct {
    Id int
    DistCost int
    Stops int
}

type MinHeap []*Node

func (h MinHeap) Len() int { return len(h)}
func (h MinHeap) Less(i, j int) bool { return h[i].DistCost < h[j].DistCost }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Node)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

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
    //h := &MinHeap{}
    h := make(MinHeap, 0, n)
    heap.Init(&h)
    heap.Push(&h, &Node{Id: src, DistCost: 0, Stops: 0})
    for h.Len() > 0 {
        city := heap.Pop(&h).(*Node)
        if city.Id == dst {
            return city.DistCost
        }
        if city.Stops == k + 1 {
            continue
        }
        dU := dist[city.Id]
        neighborEdges := graph[city.Id]
        for _, edge := range neighborEdges {
            dV, wUV := dist[edge.DstId], edge.Weight
            addToHeapCond := dU+wUV < dV || stops[edge.DstId] > city.Stops +1 
            //if (dU + wUV < dV) //minimize cost distance
            //if stops < stops[dstId] // minimize number of stops
            dist[edge.DstId] = min(dist[edge.DstId], dU+wUV)
            if addToHeapCond {
                heap.Push(&h, &Node{
                        Id: edge.DstId,
                        DistCost: city.DistCost+ edge.Weight,
                        Stops: city.Stops +1,
                    })
            }
            stops[edge.DstId] = city.Stops + 1 
        }
    }
    return -1
}




