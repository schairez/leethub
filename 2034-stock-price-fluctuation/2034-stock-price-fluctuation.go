
// k corrections for each stale max/min heap value pop operation we do log(n) operations
// time: O(klogn)
//space: O(n)

type StockPrice struct {
    minHeap *RecordHeap
    maxHeap *RecordHeap
    currTimestamp int
    timePriceMap map[int]int
}

func Constructor() StockPrice {
    minHeap := NewRecordHeap(func(a, b int) bool {
        return a < b
    })
    maxHeap := NewRecordHeap(func(a, b int) bool {
        return a > b
    })
    return StockPrice {
        minHeap: minHeap,
        maxHeap: maxHeap,
        timePriceMap: make(map[int]int),
        currTimestamp: 0, //constraint: 1 <= timestamp
    }
}


func (this *StockPrice) Update(timestamp int, price int)  {
    this.timePriceMap[timestamp] = price
    heap.Push(this.minHeap, record{timestamp, price})
    heap.Push(this.maxHeap, record{timestamp, price})
    if this.currTimestamp < timestamp {
        this.currTimestamp = timestamp
    }
}


func (this *StockPrice) Current() int {
    return this.timePriceMap[this.currTimestamp]
}


func (this *StockPrice) Maximum() int {
    for this.maxHeap.Len() != 0 {
        price, timestamp := this.maxHeap.Peek().price, this.maxHeap.Peek().timestamp 
        if currPrice := this.timePriceMap[timestamp]; currPrice == price {
            return price
        }
        heap.Pop(this.maxHeap)
    }
    return -1
}


func (this *StockPrice) Minimum() int {
    for this.minHeap.Len() != 0 {
        price, timestamp := this.minHeap.Peek().price, this.minHeap.Peek().timestamp 
        if currPrice := this.timePriceMap[timestamp]; currPrice == price {
            return price
        }
        heap.Pop(this.minHeap)
    }
    return -1
}
    


/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */

type record struct {
    timestamp int
    price int
}

type RecordHeap struct {
    records []record
    lessFn func(i, j int) bool
}

func NewRecordHeap(lessFn func(a, b int) bool) *RecordHeap {
    return &RecordHeap{lessFn : lessFn}
}
func (h *RecordHeap) Len() int { return len(h.records)}
func (h *RecordHeap) Less(i, j int) bool { return h.lessFn(h.records[i].price, h.records[j].price) }
func (h *RecordHeap) Swap(i, j int) { h.records[i], h.records[j] = h.records[j], h.records[i] }
func (h *RecordHeap) Push(v interface{} ) {
    h.records = append(h.records, v.(record))
}
func (h *RecordHeap) Pop() interface{} {
    n := h.Len()
    record := h.records[n-1]
    h.records = h.records[:n-1]
    return record
}
 
func (h *RecordHeap) Peek() record {
    return h.records[0]
}


