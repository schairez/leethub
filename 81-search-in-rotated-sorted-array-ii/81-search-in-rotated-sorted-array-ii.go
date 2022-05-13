// Input: nums = [2,5,6,0,0,1,2], target = 0
func search(nums []int, target int) bool {
    n := len(nums)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        for lo < hi && nums[lo] == nums[lo+1] && nums[lo] != target {
            lo++
        }
        for lo < hi && nums[hi] == nums[hi-1] && nums[hi] != target {
            hi--
        }
        mid := lo + (hi-lo) >> 1
        fmt.Println(mid)
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
            if nums[mid] <= target && target <= nums[hi] {
                lo = mid
            } else {
                hi = mid
            }
        }
    }
    fmt.Println(lo, hi)
    if nums[lo] == target {
        return true
    }
    if nums[hi] == target {
        return true
    }
    return false
}