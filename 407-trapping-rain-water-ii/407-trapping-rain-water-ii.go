
// 407. Trapping Rain Water II

// minHeap as pq approach
// O(m*n) space for visited map and O(m*n) space for minHeap nodes
// time: O(m*n)*log(m*n)
//

func min(a, b int) int { if a <= b { return a}; return b}
func max(a, b int) int { if a >= b { return a}; return b}

func trapRainWater(heightMap [][]int) int {
    volume := 0
    pq := NewMinHeap(func(a, b int) bool {
        return a < b
    })
    numX, numY := len(heightMap), len(heightMap[0])
    visited := make([][]bool, numX)
    for x := range visited {
        visited[x] = make([]bool, numY)
    }
    heap.Init(pq)
    for x := 0; x < numX; x++ {
        heap.Push(pq, nodeData{x, 0, heightMap[x][0] })
        visited[x][0] = true
        heap.Push(pq, nodeData{x, numY-1, heightMap[x][numY-1] })
        visited[x][numY-1] = true
    }
    for y := 1; y <= numY-2; y++ {
        heap.Push(pq, nodeData{0, y, heightMap[0][y] })
        visited[0][y] = true
        heap.Push(pq, nodeData{numX-1, y, heightMap[numX-1][y] })
        visited[numX-1][y] = true
    }
    dirs := [5]int{1, 0, -1, 0, 1}
    for pq.Len() != 0 {
        node := heap.Pop(pq).(nodeData)
        x, y := node.x, node.y
        for i := 1; i < len(dirs); i++ {
            newX, newY :=  x + dirs[i-1], y + dirs[i]
            if newX < 0 || newX >= numX ||
            newY < 0 || newY >= numY || visited[newX][newY] {
                continue
            }
            volume += max(0, node.height - heightMap[newX][newY])
            newNode := nodeData{newX, newY, max(node.height, heightMap[newX][newY]) }
            heap.Push(pq, newNode)
            visited[newX][newY] = true
        }
    }
    return volume
}

type nodeData struct {
    x      int
    y      int
    height int
}
type MinHeap struct {
    data []nodeData
    lessFn func(a, b int) bool
}

func NewMinHeap(lessFn func(i, j int) bool) *MinHeap {
    h := &MinHeap{
        lessFn : lessFn,
        data   : make([]nodeData, 0),
    }
    return h
}

// Len, Less, Swap, Push, Pop

func (h *MinHeap) Len() int {return len(h.data)}

func (h *MinHeap) Less(i, j int) bool {
    return h.lessFn(h.data[i].height, h.data[j].height)
}

func (h *MinHeap) Swap(i, j int) {
    h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(v interface{}) {
    node := v.(nodeData)
    h.data = append(h.data, node)
}
func (h *MinHeap) Pop() interface{} {
    n := len(h.data)
    item := h.data[n-1]
    h.data = h.data[:n-1]
    return item
} 

