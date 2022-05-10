func searchInsert(nums []int, target int) int {
    n := len(nums)
    lo, hi := 0, n
    for lo < hi {
        mid := lo + (hi - lo) >> 1
        if nums[mid] < target {
            lo = mid+1
        } else {
            hi = mid
        }
    }
    return lo
}