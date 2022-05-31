func sortArray(nums []int) []int {
    mergeSort(nums, 0, len(nums))
    return nums
}


func mergeSort(nums []int, left, right int) {
    if left+1 == right {
        return
    }
    mid := left + (right - left) >> 1
    mergeSort(nums, left, mid)
    mergeSort(nums, mid, right)
    
    res := make([]int, 0, right-left+1)
    x, y := left, mid
    for x < mid || y < right {
        if x == mid || y < right && nums[y] < nums[x] {
            res = append(res, nums[y])
            y++
        } else {
            res = append(res, nums[x])
            x++
        }
    }
    for i := left; i < right; i++ {
        nums[i] = res[i-left]
    }
}