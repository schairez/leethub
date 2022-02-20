// time: O(nlogk)
// space: O(k)

func kClosest(points [][]int, k int) [][]int {
    //distance from origin
    distFn := euclidDistance(0,0)
    maxHeap := NewPointsHeap(func(a, b float64) bool {
        return a > b
    })
    heap.Init(maxHeap)
    for _, point := range points {
        node := heapNode{x: point[0], y: point[1]}
        node.dist = distFn(node.x, node.y)
        if maxHeap.Len() < k {
            heap.Push(maxHeap, node)
        } else if maxHeap.Peek().dist > node.dist {
            heap.Pop(maxHeap)
            heap.Push(maxHeap, node)
        }
    }
    res := make([][]int, 0, k)
    for maxHeap.Len() != 0 {
        node := heap.Pop(maxHeap).(heapNode)
        res = append(res, []int{node.x, node.y})
    }
    return res
}
    
func euclidDistance(x1, y1 int) func(x2, y2 int) float64 {
    return func(x2, y2 int) float64 {
        return math.Sqrt(math.Pow(float64(x1-x2), 2.0) + math.Pow(float64(y1-y2), 2.0))
    }
}

type heapNode struct {
    x int
    y int
    dist float64
}
func NewPointsHeap(lessFn func(a, b float64) bool) *PointsHeap {
    return &PointsHeap{
        lessFn: lessFn,
    }
}

type PointsHeap struct {
    pointsData []heapNode
    lessFn func(i, j float64) bool
}
func(h *PointsHeap) Len() int { return len(h.pointsData)}
func(h *PointsHeap) Less(i, j int) bool {
    return h.lessFn(h.pointsData[i].dist, h.pointsData[j].dist)
}
func(h *PointsHeap) Swap(i, j int) { 
    h.pointsData[i], h.pointsData[j] = h.pointsData[j], h.pointsData[i] 
}
func(h *PointsHeap) Push( v interface{}) {
    heapNode := v.(heapNode)
    h.pointsData = append(h.pointsData, heapNode) 
}

func(h *PointsHeap) Pop() interface{} {
    n := len(h.pointsData)
    node := h.pointsData[n-1]
    h.pointsData = h.pointsData[:n-1]
    return node
}

func (h *PointsHeap) Peek() heapNode {
    return h.pointsData[0]
}

