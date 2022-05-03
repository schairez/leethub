// 34. Find First and Last Position of Element in Sorted Array
// time: O(2logn) ≈ O(logn)
// space: O(1)
// [5,7,7,8,8,10], target = 8

func searchRange(nums []int, target int) []int {
    n := len(nums)
    if n == 0 {
        return []int{-1, -1}
    }
    res := make([]int, 2)
    // search for leftmost
    start, end := 0, n-1
    for start + 1 < end {
        mid := start + (end - start) >> 1
        if nums[mid] > target {
            end = mid
        } else if nums[mid] == target {
            end = mid
        }  else {
            start = mid
        }
    }
    if nums[start] == target {
        res[0] = start
    } else if nums[end] == target {
        res[0] = end
    } else {
        res[0], res[1] = -1, -1
        return res
    }
    // search rightmost
    start, end = 0, n-1
    for start + 1 < end {
        mid := start + (end - start) >> 1
        if nums[mid] > target {
            end = mid
        } else if nums[mid] == target {
            start = mid
        } else {
            start = mid
        }
    }
    if nums[end] == target {
        res[1] = end
    } else if nums[start] == target {
        res[1] = start
    } else {
        res[0], res[1] = -1, -1
    }
        return res
}