// time: O(logn) for peekMax and popMax
// time: O(1) for top and pop operations
// space: O(n)

type MaxStack struct {
    maxHeap  *RecordHeap
    delNodeSet map[int]bool
    stack []record
    id int
}


func Constructor() MaxStack {
    maxHeap := NewRecordHeap(func(i, j int) bool {
        return i >= j
    })
    return MaxStack {maxHeap, make(map[int]bool),[]record{}, 0}
}


func (this *MaxStack) Push(x int)  {
    v := record{this.id, x}
    heap.Push(this.maxHeap, v)
    this.stack = append(this.stack, v)
    this.id++
}


func (this *MaxStack) Pop() int {
    for this.delNodeSet[this.stack[len(this.stack)-1].id] {
        this.stack = this.stack[:len(this.stack)-1]
    }
    v := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    this.delNodeSet[v.id] = true
    return v.val
}


// O(1)
func (this *MaxStack) Top() int {
    for this.delNodeSet[this.stack[len(this.stack)-1].id] {
        this.stack = this.stack[:len(this.stack)-1]
    }
    return this.stack[len(this.stack)-1].val
}

//time: O(logn)

func (this *MaxStack) PeekMax() int {
    v := this.maxHeap.Peek()
    for this.delNodeSet[v.id] { //O(1)
        heap.Pop(this.maxHeap) // O(logn)
        v = this.maxHeap.Peek()
    }
    return v.val
}

// time: O(logn)
func (this *MaxStack) PopMax() int {
    v := this.maxHeap.Peek()
    for this.delNodeSet[v.id] {
        heap.Pop(this.maxHeap)
        v = this.maxHeap.Peek()
    }
    this.delNodeSet[v.id] = true
    return v.val
}

type record struct {
    id, val int
}

func NewRecordHeap(lessFn func(i, j int) bool) *RecordHeap {
    return &RecordHeap{lessFn: lessFn, records: []record{}}
}

type RecordHeap struct {
    records []record
    lessFn func(i, j int) bool
}

//Len, Less, Swap, Push, Pop

func (h *RecordHeap) Len() int { return len(h.records)}
func (h *RecordHeap) Less(i, j int) bool {
    return h.lessFn(h.records[i].val, h.records[j].val)
}
func (h *RecordHeap) Swap(i, j int) {
    h.records[i], h.records[j] = h.records[j], h.records[i]
}
func (h *RecordHeap) Push(v any) {
    h.records = append(h.records, v.(record))
} 
func (h *RecordHeap) Pop() any {
    n := h.Len()
    v := h.records[n-1]
    h.records = h.records[:n-1]
    return v
}

func (h *RecordHeap) Peek() record {
    return h.records[0]
}



/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */