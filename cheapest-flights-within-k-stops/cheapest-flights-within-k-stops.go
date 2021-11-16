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
    for i := range flights {
        srcId := flights[i][0]
        if _, ok := graph[srcId]; !ok {
            graph[srcId] = make([]Edge, 0)
        }
        dstId, weight := flights[i][1], flights[i][2]
        graph[srcId] = append(graph[srcId], Edge{dstId, weight})
    }
    
    //minimize number of stops, s.t. K
    numStops := make([]int, n)
    //minimize distance
    //dV = dU + wUV
    dist := make([]int, n)
    for i:=0; i < n; i++ {
        dist[i] = math.MaxInt64
        numStops[i] = math.MaxInt64
    }
    dist[src] = 0
    numStops[src] = 0
    
    fmt.Println(graph)
    pq := make(PQ, 0, n)
    heap.Push(&pq, &Node{Id: src, PriceSoFar: 0, Stops: 0})
    for pq.Len() > 0 {
        fmt.Println(pq)
        city := heap.Pop(&pq).(*Node)
        fmt.Println(city.Id)
        if city.Id == dst {
            return city.PriceSoFar
        }
        if city.Stops == k + 1 {
            continue
        }
        dU := dist[city.Id]
        stops := city.Stops
        neighborEdges := graph[city.Id]
        for _, edge := range neighborEdges {
            dV, wUV := dist[edge.DstId], edge.Weight
            addToHeap := dU+wUV < dV || stops < numStops[edge.DstId] 
            //if (dU + wUV < dV) //minimize cost distance
            //if stops < numStops[dstId] // minimize number of stops
            dist[edge.DstId] = min(dist[edge.DstId], dU+wUV)
            if addToHeap {
                heap.Push(&pq, &Node{
                        Id: edge.DstId,
                        PriceSoFar: city.PriceSoFar+ edge.Weight,
                        Stops: city.Stops +1,
                    })
            }
            numStops[edge.DstId] = stops
        }
    }
    return -1
}

type Node struct {
    Id int
    PriceSoFar int
    Stops int
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq)}

func (pq PQ) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less (i, j int) bool {
    return pq[i].PriceSoFar < pq[j].PriceSoFar
}

func (pq *PQ) Push(x interface{}) {
    tmp := x.(*Node)
    *pq = append(*pq, tmp)
}


func (pq *PQ) Pop() interface{} {
    tmp := (*pq)[len(*pq)-1]
    *pq = (*pq)[0 : len(*pq)-1]
	return tmp
}



