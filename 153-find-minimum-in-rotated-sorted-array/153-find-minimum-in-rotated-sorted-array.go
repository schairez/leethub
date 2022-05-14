func findMin(nums []int) int {
    n := len(nums)
    lo, hi := 0, n-1
    // precondition
    // rotate n number times
    if nums[lo] < nums[hi] {
        return nums[lo]
    }
    for lo+1 < hi {
        mid := lo + (hi-lo) >> 1
        if nums[mid] > nums[hi] {
            lo = mid
        } else {
            hi = mid
        }
    }
    // postcondition
    if nums[lo] < nums[hi] {
        return nums[lo]
    }
    return nums[hi]
}