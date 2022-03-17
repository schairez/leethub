
type NumArray struct {
    root *SegmentNode
}


func Constructor(nums []int) NumArray {
    return NumArray{root: NewSegTree(nums)}
}


func (this *NumArray) Update(index int, val int)  {
    this.root.Update(index, val)
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.root.Query(left, right)
}

func min(a, b int) int { if a <= b { return a}; return b}
func max(a, b int) int { if a >= b { return a}; return b}


type SegmentNode struct {
    sIdx, eIdx, val int
    lChild, rChild *SegmentNode
}

func NewSegTree(nums []int) *SegmentNode {
    n := len(nums)
    var build func(int, int) *SegmentNode
    
    build = func(sIdx, eIdx int) *SegmentNode {
        if sIdx < 0 && eIdx >= n {
            return nil
        }
        if sIdx == eIdx {
            return &SegmentNode{sIdx:sIdx, eIdx:eIdx, val: nums[sIdx]}
        }
        mid := sIdx + (eIdx - sIdx) >> 1
        node := &SegmentNode{sIdx: sIdx, eIdx: eIdx}
        node.lChild = build(sIdx, mid)
        node.rChild = build(mid+1, eIdx)
        node.val = node.lChild.val + node.rChild.val
        return node
    }
    
    return build(0, n-1)
}
// [left, right] closed interval
func (seg *SegmentNode) Query(left, right int) int {
    var query func(*SegmentNode, int, int) int
    
    query = func(node *SegmentNode, sIdx, eIdx int) int {
        if node == nil || sIdx > eIdx || 
        node.sIdx > eIdx || node.eIdx < sIdx {
            return 0
        }
        if node.sIdx == sIdx && node.eIdx == eIdx {
            return node.val
        }
        mid := node.sIdx + (node.eIdx - node.sIdx) >> 1
        sum := 0
        sum += query(node.lChild, sIdx, min(eIdx,mid))
        sum += query(node.rChild, max(mid+1, sIdx), eIdx)
        return sum
    }
    return query(seg, left, right)
}


func (seg *SegmentNode) Update(idx, newVal int) {
    var update func(node *SegmentNode)
    update = func(node *SegmentNode) {
        if node == nil {
            return
        }
        if node.sIdx == idx && node.eIdx == idx {
            node.val = newVal
            return
        }
        l, r := node.sIdx, node.eIdx
        mid := l + (r-l) >> 1
        if idx > mid {
            update(node.rChild)
        } else {
            update(node.lChild)
        }
        node.val = node.lChild.val + node.rChild.val
    }
    
    update(seg)
}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */