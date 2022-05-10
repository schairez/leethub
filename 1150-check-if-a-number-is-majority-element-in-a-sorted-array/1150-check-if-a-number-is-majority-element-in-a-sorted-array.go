// 1150. Check If a Number Is Majority Element in a Sorted Array
// time: O(2logn) ≈ O(logn)
// space: O(1)
func isMajorityElement(nums []int, target int) bool {
    // find leftmost idx of t
    n := len(nums)
    l, r := 0, n-1
    for l + 1 < r {
        mid := l + (r-l) >> 1
        if nums[mid] >= target {
            r = mid
        } else {
            l = mid
        }
    }
    var leftMost int
    if nums[l] == target {
        leftMost = l
    } else if nums[r] == target {
        leftMost = r
    } else {
        return false
    }
    l, r = 0, n-1
    for l + 1 < r {
        mid := l + (r-l) >> 1
        if nums[mid] <= target {
            l = mid
        } else {
            r = mid
        }
    }
    var rightMost int
    if nums[r] == target {
        rightMost = r
    } else if nums[l] == target {
        rightMost = l
    }
    return (rightMost - leftMost + 1) > n / 2
    
}