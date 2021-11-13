func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    tmp := 1
    for i := range nums {
        res[i] = tmp
        tmp = tmp*nums[i]
    }
    tmp = nums[len(nums)-1]
    for i:=len(nums)-2; i >=0; i-- {
        res[i] *= tmp
        tmp = tmp*nums[i]
    }
    return res
}


/*

              [1,1,1,1]
Input: nums = [1,2,3,4]
            tmp:=1
            for i := range nums {
                res[i] = tmp
                tmp = tmp*arr[i]
            }




Output: [24,12,8,6]




*/