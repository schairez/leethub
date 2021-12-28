//prefix sum approach
//space: O(n)
//time: O(1) for query, for construction: O(n)

type NumArray struct {
    prefixSumArr []int
}


func Constructor(nums []int) NumArray {
    n := len(nums)
    runningSum := 0
    prefixSumArr := make([]int, n+1)
    prefixSumArr[0] = 0
    for idx, num := range nums {
        runningSum += num
        prefixSumArr[idx+1] = runningSum
    }
    
    return NumArray{ prefixSumArr }
}

//[-2, 0, 3, -5, 2, -1]
//[0, -2, -2, 1, -4, -2, -3]
//f(0,0) = prefix[1] - prefix[0]
//f(0,2) = prefix[3] - prefix[0]
//constraint:
//0 <= left <= right < nums.length

func (this *NumArray) SumRange(left int, right int) int {
    return this.prefixSumArr[right+1] - this.prefixSumArr[left]
}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */


/*
segment Tree approach; overkill for an immutable question xD
time: O(n) for constjuctor
time: O(logn) for sumQuery 

space: O(n) for query construction
space: O(logn) for stack space in range query search

*/

type NumArrayV2 struct {
    root *SegmentNode
    nums []int
    
}

func ConstructorV2(nums []int) NumArrayV2 {
    query := &NumArrayV2{nums: nums}
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
func (this *NumArrayV2) buildTree(startIdx, endIdx int) *SegmentNode {
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

func (this *NumArrayV2) rangeSum(node *SegmentNode, left int, right int) int {
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


func (this *NumArrayV2) SumRange(left int, right int) int {
    return this.rangeSum(this.root, left, right)
}


