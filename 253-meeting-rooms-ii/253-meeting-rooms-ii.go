//time: sort takes O(nlogn) and we iterate through
// n interval values and perform a push or pop operation which amortizes to O(logn) time
// thus, time: O(nlogn)
// space: sort takes O(logn) space and heap uses O(n) space worst case, thus O(n) 


// sort by start time
// push the endTime value to heap
// when comparing the next iter j startTime with the top i of minHeap 
// if startTimej > endTimei, let's pop from our heap 
// intuitively, this new meeting schedule can be added to the prev occupied room
// the len(minHeap) will tell us the number of roomsHeap we require 
func minMeetingRooms(intervals [][]int) int {
    
    if len(intervals) < 2 {
        return len(intervals)
    }
    sort.Slice(intervals, func(i,j int) bool{
        return intervals[i][0] < intervals[j][0]
    })
    
    roomsHeap := make(minHeap, 0)
    heap.Init(&roomsHeap)
    for idx := 0; idx < len(intervals); idx++ {
        if roomsHeap.Len() > 0 && intervals[idx][0] >= roomsHeap.Peek() {
            heap.Pop(&roomsHeap)
        }
        heap.Push(&roomsHeap, intervals[idx][1]) 
    }
    return roomsHeap.Len()
}

type minHeap []int 

func (h minHeap) Peek() int {
    return h[0]
}

func (h minHeap) Len() int{
    return len(h)
}

func (h minHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v interface{}){
    *h = append(*h, v.(int))
} 
func (h *minHeap) Pop() (interface{}) {
    old := *h
    n := len(*h)
    x := old[n-1]
    *h = old[:n-1]
    return x
}