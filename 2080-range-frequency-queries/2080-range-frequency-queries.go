
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








/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */