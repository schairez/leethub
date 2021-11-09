/*
adding each node is a log(n) operation but for n nodes so time: O(nlogn)
space: O(1)

*/

import "fmt"



func findKthLargest(nums []int, k int) int {
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
    fmt.Println(nums)
    return nums[n-k]
    
}
