//153. Find Minimum in Rotated Sorted Array
//time:O(logN)
//space:O(1)

func findMin(nums []int) int {
    n := len(nums)
    if nums[0] < nums[n-1] {
        return nums[0]
    }
    lo, hi := 0, n-1
    for lo + 1 < hi {
        mid := lo + (hi-lo) >> 1
        if nums[mid] > nums[hi] || nums[mid] > nums[lo] {
            lo = mid
        } else {
            hi = mid
        }
    }
    // constraint: nums are unique
    if nums[lo] < nums[hi] {
        return nums[lo]
    }
    return nums[hi]
}