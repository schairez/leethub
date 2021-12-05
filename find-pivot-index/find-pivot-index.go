//time : O(2n) ~ O(n)
//space: O(1)

func pivotIndex(nums []int) int {
    sumVals := func() int {
        v := 0
        for i := range nums {
            v += nums[i]
        }
        return v
    }
    res := -1
    totalSum := sumVals() 
    leftSum := 0
    for idx := range nums {
        num := nums[idx]
        rightSum := totalSum - num - leftSum
        if leftSum == rightSum { return idx }
        leftSum += num
    }
    return res
}