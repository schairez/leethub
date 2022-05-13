
//81. Search in Rotated Sorted Array II
// time: avg: O(logn) worst: O(n)
// space: O(1)
// Input: nums = [2,5,6,0,0,1,2], target = 0
func search(nums []int, target int) bool {
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
        if nums[mid] == target {
            return true
        }
        if nums[lo] <= nums[mid] {
            // nums b/w lo and mid are sorted
            if nums[lo] <= target && target <= nums[mid] {
                hi = mid
            } else {
                lo = mid
            }
        } else {
            // nums b/w mid and end are sorted
            if nums[mid] <= target && target <= nums[hi] {
                lo = mid
            } else {
                hi = mid
            }
        }
    }
    if nums[lo] == target {
        return true
    }
    if nums[hi] == target {
        return true
    }
    return false
}