func findPeakElement(nums []int) int {
    n := len(nums)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if nums[mid] > nums[mid+1] {
            hi = mid
        } else {
            lo = mid
        }
    }
    if n == 1 || nums[lo] > nums[lo+1] {
        return lo
    }
    return hi
}