
func findMin(nums []int) int {
    n := len(nums)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        for lo < hi && nums[lo] == nums[lo+1] {
            lo++
        }
        for lo < hi && nums[hi] == nums[hi-1] {
            hi--
        }
        mid := lo + (hi-lo) >> 1
        if nums[mid] <= nums[hi] {
            hi = mid
        } else {
            lo = mid
        }
    }
    if nums[lo] < nums[hi] {
        return nums[lo]
    }
    return nums[hi]
}