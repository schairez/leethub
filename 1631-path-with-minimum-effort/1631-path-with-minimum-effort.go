
func minimumEffortPath(heights [][]int) int {
    numR, numC := len(heights), len(heights[0])
    maxInt32 := 1 << 31 - 1
    dist := make([][]int, numR)
    for x := range dist {
        dist[x] = make([]int, numC)
        for y := 0; y < numC; y++ {
            dist[x][y] = maxInt32 
        }
    }
    inArea := func(x, y int) bool {
        return x >= 0 && x < numR && y >= 0 && y < numC
    }
    dirs := [5]int{1,0,-1,0,1}
    pq := newPQ(func(a, b int) bool {
        return a < b
    })
    heap.Init(pq)
    heap.Push(pq, nodeData{0, 0, 0})
    for pq.Len() != 0 {
        data := heap.Pop(pq).(nodeData)
        if data.x == numR-1 && data.y == numC-1 {
            return data.diff
        } 
        for i := 1; i < 5; i++ {
            newX, newY := data.x + dirs[i-1], data.y + dirs[i]
            if !inArea(newX, newY) {
                continue
            }
            candEffort := max(data.diff, abs(heights[data.x][data.y] - heights[newX][newY]))
            if candEffort < dist[newX][newY] {
                data := nodeData{newX, newY, candEffort}
                heap.Push(pq, data)
                dist[newX][newY] = candEffort
            }
        }
    }
    return -1
}

func max(a, b int) int { if a >= b { return a}; return b}

type nodeData struct {
    x, y, diff int
}

type PriorityQueue struct {
    data []nodeData
    lessFn func(i, j int) bool
}

func newPQ(lessFn func(i, j int) bool) *PriorityQueue {
    return &PriorityQueue{
        lessFn: lessFn,
        data: make([]nodeData, 0),
    }
}

func (pq *PriorityQueue) Len() int {
    return len(pq.data)
}
func (pq *PriorityQueue) Less(i, j int) bool {
    return pq.lessFn(pq.data[i].diff, pq.data[j].diff)
} 
func (pq *PriorityQueue) Swap(i, j int) {
    pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}
func (pq *PriorityQueue) Push(v interface{}) {
    pq.data = append(pq.data, v.(nodeData))
}
func (pq *PriorityQueue) Pop() interface{} {
    n := pq.Len()
    v := pq.data[n-1]
    pq.data = pq.data[:n-1]
    return v
}


func abs(x int) int {if x < 0 { return -x}; return x}

