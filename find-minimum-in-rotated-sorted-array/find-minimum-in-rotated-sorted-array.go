//time:O(logN)
//space:O(1)


func findMin(nums []int) int {
    var mid int
    lo, hi := 0, len(nums)-1
    for nums[lo] > nums[hi] {
        mid = lo + (hi-lo) >> 1
        if nums[mid] < nums[hi] {
            hi = mid
        } else {
            lo = mid+1
        }
        
    }
    return nums[lo]
    
}

/*
        if nums[mid] >= nums[lo] {
            lo = mid+1
        } else {
            hi = mid
        }
*/