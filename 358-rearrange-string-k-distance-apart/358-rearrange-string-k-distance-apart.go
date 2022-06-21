
func rearrangeString(s string, k int) string {
    n := len(s)
    var freqTbl [26]int 
    for idx := range s {
        freqTbl[s[idx]-'a']++
    }
    pq := &Heap{data: make([]*Pair, 0, n)}
    pq.lessFn = func(a, b int) bool {
        return a > b
    }
    heap.Init(pq)
    for idx := range freqTbl {
        if freqTbl[idx] != 0 {
            heap.Push(pq, &Pair{freq: freqTbl[idx], char: idx })
        }
    }
    var node *Pair
    queue := make([]*Pair, 0, k)
    var sb strings.Builder
    sb.Grow(n)
    for pq.Len() != 0 {
        node = heap.Pop(pq).(*Pair)
        sb.WriteByte(byte(node.char+'a'))
        node.freq--
        queue = append(queue, node)
        if len(queue) >= k {
            node, queue = queue[0], queue[1:]
            if node.freq > 0 {
                heap.Push(pq, node)
            }
        }
    }
    kDistStr := sb.String()
    if len(kDistStr) == n {
        return kDistStr
    }
    return ""
}


type Pair struct {
    freq int
    char int
}

type Heap struct {
    data []*Pair
    lessFn func(i, j int) bool
}

func (h *Heap) Len() int {
    return len(h.data)
}

func (h *Heap) Less(i, j int) bool {
    return h.lessFn(h.data[i].freq, h.data[j].freq)
}
func (h *Heap) Swap(i, j int) {
    h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *Heap) Push(v interface{}) {
    h.data = append(h.data, v.(*Pair))
}
func (h *Heap) Pop() interface{} {
    n := h.Len()
    v := h.data[n-1]
    h.data[n-1] = nil
    h.data = h.data[:n-1]
    return v
}



