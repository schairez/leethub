// 2161. Partition Array According to Given Pivot
// time: O(n)
// space: O(n)

func pivotArray(nums []int, pivot int) []int {
    n := len(nums)
    left, mid, right := 0, 0, n 
    for _, v := range nums {
        if v < pivot {
            mid++
        } else if v > pivot {
            right--
        }
    }
    res := make([]int, n)
    for _, v := range nums {
        if v < pivot {
            res[left] = v
            left++
        } else if v == pivot {
            res[mid] = v
            mid++
        } else {
            res[right] = v
            right++
        }
    }
    return res
}