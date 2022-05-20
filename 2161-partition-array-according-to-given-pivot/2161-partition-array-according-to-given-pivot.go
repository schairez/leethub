func pivotArray(nums []int, pivot int) []int {
    var less, greater, equal []int
    for _, v := range nums {
        if v < pivot {
            less = append(less, v)
        } else if v > pivot {
            greater = append(greater, v)
        } else {
            equal = append(equal, v)
        }
    }
    n3 := len(less) + len(equal) + len(greater)
    res := make([]int, 0, n3)
    for i := range less {
        res = append(res, less[i]) 
    }
    for i := range equal {
        res = append(res, equal[i]) 
    }
    for i := range greater {
        res = append(res, greater[i]) 
    }
    return res
}