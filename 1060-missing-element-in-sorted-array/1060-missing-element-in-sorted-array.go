func missingElement(nums []int, k int) int {
    start, end := 0, len(nums)-1
    var numMissing int
    var mid int
    for start <= end {
        mid = start + (end-start) >> 1
        numMissing = nums[mid] - (nums[0]+ mid)
         if numMissing < k { // [start, m]
            start = mid+1
         } else {
             end = mid-1
         }
    }
    numMissing = nums[end] - (nums[0]+end)
    return nums[end] + (k - numMissing)
    
}

        //numMissing = numElemsInRange - numVals
        //to the left num
        // num elems in range
        //numElemsInRange := nums[mid] - nums[start] + 1 
        // numVals in array from left to right 
        // numVals := mid - start + 1