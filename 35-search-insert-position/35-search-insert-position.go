

func searchInsert(nums []int, target int) int {
    n := len(nums)
    lo, hi := 0, n
    for lo < hi {
        mid := lo + (hi - lo) >> 1
        if nums[mid] >= target {
            hi = mid
        } else {
            lo = mid+1
        }
    }
    return lo
}
    

func searchInsertV3(nums []int, target int) int {
    n := len(nums)
    lo, hi := 0, n-1
    for lo + 1 < hi {
        mid := lo + (hi - lo) >> 1
        if nums[mid] <= target {
            lo = mid
        } else {
            hi = mid
        }
    }
    if nums[lo] >= target {
        return lo
    }
    if nums[hi] >= target {
        return hi
    }
    if nums[hi] < target {
        return hi+1
    }
    return lo
}


func searchInsertV2(nums []int, target int) int {
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