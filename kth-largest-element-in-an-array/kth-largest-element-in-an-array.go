//time: O(NlogK)
//space: O(K)

func findKthLargest(nums []int, k int) int {
    h := &PriorityQueue{ []int{} }
    heap.Init(h)
    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    return heap.Pop(h).(int)
}

//minHeap based pq
type PriorityQueue struct {
    items []int
}
//Pop, Push, Swap, Len, Less for heap interface impl
func (h *PriorityQueue) Len() int { return len(h.items) }
func (h *PriorityQueue) Less(i, j int) bool { 
    return h.items[i] < h.items[j] 
}
func (h *PriorityQueue) Swap(i, j int) {
    h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h *PriorityQueue) Push(v interface{}) {
    h.items = append(h.items, v.(int) )
}
func (h *PriorityQueue) Pop() interface{} {
    n := len(h.items)
    x := h.items[n-1]
    h.items = h.items[:n-1]
    return x
}

/////////////////////////////////////////////////////////////////////////////////
/*
version 11/8/2021
adding each node is a log(n) operation but for n nodes so time: O(nlogn)
space: O(1)
*/

func findKthLargest2(nums []int, k int) int {
    n := len(nums)
    
    parent := func(i int) int { return (i >> 1) - 1  }
    leftChild := func(i int) int { return (i << 1) + 1 }
    rightChild := func(i int) int {  return (i << 1) + 2  }
    swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
    
    var heapify func(parIdx, n int)
    //heapify; builds maxHeap and checks for maxHeap property ( vals at child nodes are < parent)
    heapify = func(parIdx, n int ) {
        lcIdx := leftChild(parIdx)
        rcIdx :=  rightChild(parIdx)
        if lcIdx >= n { return }
        largestChild := lcIdx
        if rcIdx < n && nums[rcIdx] > nums[lcIdx] {
            largestChild = rcIdx
        } 
        if nums[largestChild] > nums[parIdx] {
            swap(largestChild, parIdx)
            heapify(largestChild, n)
        }
    }
    //open-closed [0,n)
    heapSort := func(n int) {
        //last non-leaf node
        p := parent(n)
        //build maxHeap
        for i:=p; i>=0; i-- {
            heapify(i, n)
        }
        i := n-1
        for i >=0 {
            swap(0, i)
            heapify(0, i-1)
            i--
        }
    }
    heapSort( n)
    return nums[n-k]
    
}
