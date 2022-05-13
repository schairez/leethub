// 33. Search in Rotated Sorted Array
// time: O(logn)
// space: O(1)
func search(nums []int, target int) int {
    n := len(nums)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if nums[lo] < nums[mid] {
            if nums[lo] <= target && target < nums[mid] {
                hi = mid
            } else {
                lo = mid
            }
        } else {
            if nums[mid] < target && target <= nums[hi] {
                lo = mid
            } else {
                hi = mid
            }
        }
    }
    if nums[lo] == target {
        return lo
    }
    if nums[hi] == target {
        return hi
    }
    return -1
}