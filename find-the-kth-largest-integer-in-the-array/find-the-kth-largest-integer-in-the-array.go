import "container/heap"

//we iterate through the nums slice in O(N) time
//heap insert and poll operations are O(logK) time
//at worst, we compare each byte index in a M-sized string num item 
//in O(M) time where m is the size of the string item 
//therefore time: O(NlogK * M) time
//space: O(K), our heap only has K -sized underlying slice len

func kthLargestNumber(nums []string, k int) string {
    h := &PriorityQueue{ []string{} }
    heap.Init(h)
    for _, val := range nums {
        heap.Push(h, val)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    return heap.Pop(h).(string)
    
}


// minHeap based approach
// we keep up to K elements in our Heap 
// when len of Heap > K; poll min node 
// return pollNode which is the Kth largest

type PriorityQueue struct {
    items []string
}
func (h *PriorityQueue) Len() int { return len(h.items) }

func (h *PriorityQueue) Less(i, j int) bool { 
    if len(h.items[i]) == len(h.items[j]) {
        return h.items[i] < h.items[j]
    }
    return len(h.items[i]) < len(h.items[j])
}

func (h *PriorityQueue) Swap(i, j int) { 
    h.items[i], h.items[j] = h.items[j], h.items[i] 
}

func (h *PriorityQueue) Push(v interface{}) { 
    h.items = append(h.items, v.(string)) 
}

func (h *PriorityQueue) Pop() interface{} {
    n := len(h.items)
    x := h.items[n-1]
    h.items = h.items[:n-1]
    return x
}