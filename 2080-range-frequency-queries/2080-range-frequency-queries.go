// 2080. Range Frequency Queries
// time: O(logn) for query; O(n) to construct
// space: O(n)


const size = 1e5 

type RangeFreqQuery struct {
    data [size+1][]int // val to indices bucket mapping
}


func Constructor(arr []int) RangeFreqQuery {
    rfq := RangeFreqQuery{}
    for idx, val := range arr {
        rfq.data[val] = append(rfq.data[val], idx)
    }
    return rfq
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
    arr := this.data[value]
    if len(arr) == 0 {
        return 0
    }
    start := leftMostIdx(arr, left)
    if start == -1 {
        return 0
    }
    end := rightMostIdx(arr, right)
    return end - start + 1
}

func rightMostIdx(arr []int, target int) int {
    n := len(arr)
    lo, hi := 0, n-1 
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if arr[mid] <= target {
            lo = mid
        } else {
            hi = mid
        }
    }
    if arr[hi] <= target {
        return hi
    } else if arr[lo] <= target {
        return lo
    }
    return -1
}

func leftMostIdx(arr []int, target int) int {
    n := len(arr)
    lo, hi := 0, n-1 
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if arr[mid] >= target {
            hi = mid
        } else {
            lo = mid
        } 
    }
    if arr[lo] >= target {
        return lo
    } else if arr[hi] >= target {
        return hi
    }
    return -1
} 




/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */


/*
// seg tree version

type RangeFreqQuery struct {
    root *SegNode
}


func Constructor(arr []int) RangeFreqQuery {
    n := len(arr)
    root := newNode(0, n-1)
    for i := 0; i < n; i++ {
        // update freq for node val and preorder propagate
        // to qualifying child nodes
        root.Insert(i, arr[i])
    }
    return RangeFreqQuery{root}
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
    return this.root.QueryRange(left, right, value)
}


type SegNode struct {
    freq map[int]int
    start, end int
    left, right *SegNode
}

func newNode(sIdx, eIdx int) *SegNode {
    segNode := &SegNode{start: sIdx, end: eIdx}
    segNode.freq = make(map[int]int, eIdx-sIdx+1)
    return segNode
}

func (seg *SegNode) Insert(sIdx int, val int) {
    seg.freq[val]++
    if seg.start == seg.end {
        return
    }
    mid := seg.start + (seg.end - seg.start) >> 1
    if sIdx <= mid {
        if seg.left == nil {
            seg.left = newNode(seg.start, mid)
        }
        seg.left.Insert(sIdx, val)
    } else {
        if seg.right == nil {
            seg.right = newNode(mid+1, seg.end)
        }
        seg.right.Insert(sIdx, val)
    }
} 
func (seg *SegNode) QueryRange(sIdx int, eIdx int, val int) int {
    // if nil or disjoint interval
    if eIdx < sIdx || seg == nil || seg.start > eIdx || seg.end < sIdx {
        return 0
    }
    if seg.start == sIdx && seg.end == eIdx {
        return seg.freq[val]
    }
    mid := seg.start + (seg.end - seg.start) >> 1
    if mid >= eIdx {
        return seg.left.QueryRange(sIdx, eIdx, val)
    } else if mid < sIdx {
        return seg.right.QueryRange(sIdx, eIdx, val)
    }
    return seg.left.QueryRange(sIdx, mid, val) + seg.right.QueryRange(mid+1, eIdx, val)
}


*/




/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

