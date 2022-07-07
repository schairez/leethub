// 871. Minimum Number of Refueling Stops
// time: O(nlogn)
// space: O(n)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
    n := len(stations)
    pq := newIntHeap(func(i, j int) bool {
        return i > j
    })
    pos, idx, numStops := 0, 0, 0
    heap.Init(pq)
    heap.Push(pq, startFuel)
    for pq.Len() != 0 {
        fuel := heap.Pop(pq).(int)
        pos += fuel 
        if pos >= target {
            return numStops 
        }
        for idx < n && pos >= stations[idx][0] {
            heap.Push(pq, stations[idx][1])
            idx++
        }
        numStops++
    }
    return -1
}

type IntHeap struct {
    lessFn func(i, j int) bool
    data []int
}
func (h *IntHeap) Len() int {
    return len(h.data)
}
func (h *IntHeap) Less(i, j int) bool {
    return h.lessFn(h.data[i], h.data[j])
}
func (h *IntHeap) Swap(i, j int) {
    h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *IntHeap) Push(v interface{}) {
    h.data = append(h.data, v.(int))
}
func (h *IntHeap) Pop() interface{} {
    n := h.Len()
    v := h.data[n-1]
    h.data = h.data[:n-1]
    return v
}

func newIntHeap(lessFn func(i, j int) bool) *IntHeap {
    h := &IntHeap{}
    h.lessFn = lessFn
    return h
}
