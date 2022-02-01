//import "container/heap"

/*
- the median is the middle pt in our sequence
  - left half of list is <= median
  - right half of list is >= median

two heaps approach
we're only interested in the middle most elems; if 
we were to split our sorted sequence into two equal arrs, 
we'd only be interested in the rightmost value of the left
arr (the maxVal) and the leftmost val of the largerValued/right arr ( the minVal).
in other words, we'd only be interested in the maxElem of the left half
and the minElem of the right half

if we use two heaps, a maxHeap and a minHeap; maxHeap = left; minHeap = right
- we store newly inserted val into maxHeap 
- we pop the maxElem from maxHeap and insert into our minHeap
- balance if maxHeap < minHeap: we poll from minHeap and push to maxHeap
- both heaps are balanced; diff in size <= 1

to find the median
- if len(maxHeap) == len(minHeap); we have even # of elems, thus we
  take the mean of the top of the heaps
- else: 
  median is @ top of maxHeap

*/


type MedianFinder struct {
    smallerNums *IntHeap
    largerNums *IntHeap
}


func Constructor() MedianFinder {
    minHeap := NewHeap(func(a, b int) bool {
        return a < b
    })
    maxHeap := NewHeap(func(a, b int) bool {
        return a > b
    })
    return MedianFinder{smallerNums : maxHeap, 
                        largerNums : minHeap}
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.smallerNums, num)
    heap.Push(this.largerNums, heap.Pop(this.smallerNums))
    if this.smallerNums.Len() < this.largerNums.Len() {
        heap.Push(this.smallerNums, heap.Pop(this.largerNums))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    n := this.smallerNums.Len() + this.largerNums.Len()
    if n % 2 == 0 {
        v1 := this.smallerNums.Peek()
        v2 := this.largerNums.Peek()
        return float64(v1 + v2) / float64(2)
    }
    return float64(this.smallerNums.Peek())
}



/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */


type IntHeap struct {
    nums []int
    lessFn func(i, j int) bool
}

func NewHeap(less func(i, j int) bool) *IntHeap {
    return &IntHeap{
        lessFn: less,
    }
}

func (h *IntHeap) Len() int { return len(h.nums)}

func (h *IntHeap) Less(i, j int) bool {
    return h.lessFn(h.nums[i], h.nums[j])
}
func (h *IntHeap) Swap(i, j int) {h.nums[i], h.nums[j] = h.nums[j], h.nums[i]}

func (h *IntHeap) Push(v interface{}) {
    val := v.(int)
    h.nums = append(h.nums, val)
}

func (h *IntHeap) Pop() interface{} {
    n := h.Len()
    val := h.nums[n-1]
    h.nums = h.nums[:n-1]
    return val
}

func (h *IntHeap) Peek() int {
    return h.nums[0]
}

