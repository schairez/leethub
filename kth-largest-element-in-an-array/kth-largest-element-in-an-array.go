/*
adding each node is a log(n) operation but for n nodes so time: O(nlogn)
space: O(n)

*/

import "fmt"

/*
root is at idx0
---
parent is at i >> 1 
parent is at nums[(i-2)/2] for (i-2)/2 >= 0
---
leftChild is at (i << 1) + 1
leftChild is at nums[i*2+1]
rightChild is at (i << 1 ) + 2 
rightChild is at nums[i*2+2]
*/

/*
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

*/



func findKthLargest(nums []int, k int) int {
    n := len(nums)
    heapSort(nums, n)
    fmt.Println(nums)
    return nums[n-k]
    
}

func leftChild(i int) int {
    return (i << 1) + 1
}

func rightChild(i int) int {
    return (i << 1 ) + 2
}

func getParent(i int) int {
        return (i >> 1) - 1 
}

func heapSort(nums []int, n int) {
    
    //we start from the middle elem
    p := getParent(n) 
    for i := p; i >=0; i -- {
        minHeapify(nums, i, n)
        
    }
    i := n-1
    for i >= 0 {
        nums[0], nums[i] = nums[i], nums[0]
        minHeapify(nums, 0, i-1)
        i--
    }
}

func minHeapify(nums []int, parent, n int) { 
    if parent < 0 { return }
    
    lc := leftChild(parent)
    rc := rightChild(parent)
    //check if lc is within limits
    if lc >= n { return }
    
    var largest int
    //assume largest is at lc
    largest = lc
    if rc < n && nums[rc] > nums[lc] {
        largest = rc 
    }     
    if nums[largest] > nums[parent] {
        //swap(largest, parent)
        nums[largest], nums[parent] = nums[parent], nums[largest]
        minHeapify(nums, largest, n)
    }
    
}


