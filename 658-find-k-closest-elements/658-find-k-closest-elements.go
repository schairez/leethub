// 658. Find K Closest Elements
// time: O(logn + k)
// space: O(1)
func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    if k == n {
        return arr
    }
    left := closestElem(arr, x)
    right := left
    for right - left + 1 < k {
        if left == 0 {
            right++
        } else if (right == n-1) {
            left--
        } else {
            if abs(x - arr[left-1]) <= abs(x - arr[right+1]) {
                left--
            } else {
                right++
            }
        }
    }
    return arr[left:right+1]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func closestElem(nums []int, target int) int {
    n := len(nums)
    lo, hi := 0, n-1
    for lo+1 < hi {
        mid := lo + (hi-lo) >> 1
        if nums[mid] >= target {
            hi = mid
        } else {
            lo = mid
        }
    }
    if nums[lo] == target || 
    abs(target-nums[lo]) <= abs(target - nums[hi]) {
        return lo
    }
    return hi
}