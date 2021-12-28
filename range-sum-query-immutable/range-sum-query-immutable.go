/*
segment Tree approach
time: O(n) for constjuctor
time: O(logn) for sumQuery 

space: O(n) for query construction
space: O(logn) for stack space in range query search

*/

type NumArray struct {
    root *SegmentNode
    nums []int
    
}

func Constructor(nums []int) NumArray {
    query := &NumArray{nums: nums}
    start, end := 0, len(nums)-1
    query.root = query.buildTree(start, end)
    return *query
}

type SegmentNode struct {
    //closed interval [startIdx, endIdx]
    startIdx, endIdx, sum int
    left, right *SegmentNode
}
//segment tree; postorder node traversal
func (this *NumArray) buildTree(startIdx, endIdx int) *SegmentNode {
    node := &SegmentNode{startIdx: startIdx, endIdx: endIdx}
    if startIdx == endIdx {
        node.sum = this.nums[startIdx]
        return node
    }
    mid := startIdx + (endIdx - startIdx) / 2
    node.left = this.buildTree(startIdx, mid)
    node.right = this.buildTree(mid+1, endIdx)
    node.sum = node.left.sum + node.right.sum
    return node
    
}
func min(a, b int) int { if a <= b { return a }; return b }
func max(a, b int) int { if a >= b { return a }; return b }

func (this *NumArray) rangeSum(node *SegmentNode, left int, right int) int {
    if node == nil || left > right || left > node.endIdx || right < node.startIdx {
        return 0
    }
    if left == node.startIdx && right == node.endIdx {
        return node.sum
    }
    mid := node.startIdx + (node.endIdx - node.startIdx) / 2
    sum := 0
    sum += this.rangeSum(node.left, left, min(mid, right))
    sum += this.rangeSum(node.right, max(mid+1, left), right)
    return sum
    
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.rangeSum(this.root, left, right)
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */