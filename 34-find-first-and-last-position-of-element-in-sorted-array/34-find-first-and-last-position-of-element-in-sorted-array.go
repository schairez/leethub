
// 34. Find First and Last Position of Element in Sorted Array
//
// time: O(2*logn) ≈ O(logn)
// space: O(1)

func searchRange(nums []int, target int) []int {
    n := len(nums)
    if n == 0 {
        return []int{-1, -1}
    }
    binSearch := func(findLeftMost bool) int {
        start, end := 0, n-1
        for start + 1 < end { // terminates: start + 1 == end
            mid := start + (end - start) >> 1
            if nums[mid] < target {
                start = mid
            } else if nums[mid] > target {
                end = mid
            } else {
                // if findLeftmost and idx == target :
                // we want to continue our search 
                if findLeftMost {
                    end = mid
                    continue
                }
                start = mid
            }
        }
        if findLeftMost {
            if nums[start] == target {
                return start
            } else if nums[end] == target {
                return end
            }
        } else {
            if nums[end] == target {
                return end
            } else if nums[start] == target {
                return start
            }
        }
        return -1
    }
    res := make([]int, 2)
    findLeft := true
    res[0] = binSearch(findLeft)
    if res[0] == -1 {
        res[1] = -1
        return res
    }
    res[1] = binSearch(!findLeft)
    if res[1] == -1 {
        return res
    }
    return res
}