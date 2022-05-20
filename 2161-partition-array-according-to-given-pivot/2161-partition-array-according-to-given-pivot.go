func pivotArray(nums []int, pivot int) []int {
    n := len(nums)
    res := make([]int, n)
    idx := 0
    for _, v := range nums {
        if v < pivot {
            res[idx] = v 
            idx++
        }
    }
    for _, v := range nums {
        if v == pivot {
            res[idx] = v 
            idx++
        }
    }
    for _, v := range nums {
        if v > pivot {
            res[idx] = v 
            idx++
        }
    }
    return res
}