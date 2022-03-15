
type NumArray struct {
    seg SegTree
}


func Constructor(nums []int) NumArray {
    return NumArray{seg: NewSegTree(nums)}
}


func (this *NumArray) Update(index int, val int)  {
    this.seg.Update(index, val)
}


// since [left,right] in this version
func (this *NumArray) SumRange(left int, right int) int {
    //right closed
    return this.seg.Query(left, right)
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */


type SegTree struct {
    tree    []int
    arrSize int
}

/*
func (seg SegTree) Build(nums []int) {
    
}
*/
      

func NewSegTree( nums []int) SegTree {
    // build segTree nodes
    n := len(nums)
    tree := make([]int, 4*n)
    for i := 0; i < n; i++ {
        tree[n+i] = nums[i]
    }
    for i := n-1; i > 0; i-- {
        tree[i] = tree[2*i] + tree[2*i+1]
        //tree[i] = tree[i<<1] + tree[i<<1 | 1]
    }
    
    return SegTree{tree, n}
}


func (seg SegTree) Update(idx, val int) {
    //fmt.Println(seg.tree)
    nodeIdx := idx + seg.arrSize
    seg.tree[nodeIdx] = val 
    
    // update parent nodes bottom up the chain
    for i := nodeIdx; i > 1; i = i >> 1  {
        seg.tree[i>>1]  = seg.tree[i] + seg.tree[i^1]
    }
}
//[left, right)
func (seg SegTree) Query(left, right int) int {
    //fmt.Println("query")
    //fmt.Println(seg.tree)
    n := seg.arrSize
    res := 0
    l, r := left + n, right + n
    for l <= r {
        if l & 1 == 1 {
            res += seg.tree[l]
            l++
        }
        if r & 1 == 0 {
            res += seg.tree[r]
            r--
        }
        l >>= 1
        r >>= 1
    }
    return res
} 
