func findPeakElement(nums []int) int {
    /// lo, hi
    start, end := 0, len(nums)-1
    //[2,1]
    for start < end {
        mid := start + (end-start) >> 1
        if nums[mid] < nums[mid+1] {
            start = mid+1
        } else {
            end = mid
        }
            // l = mid+1
    }
    return start
    
}
/*
m = 2 + (3-2)/2 = 2
            l
n ums = [1,2,3,1]
          |
          m
*/