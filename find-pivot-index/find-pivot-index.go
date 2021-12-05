//time : O(2n) ~ O(n)
//space: O(1)

func pivotIndex(nums []int) int {
    totalSum := func() int {
        v := 0
        for i := range nums {
            v += nums[i]
        }
        return v
    }()
    
    res := -1
    leftSum := 0
    //rightSum := totalSum - num - leftSum
    for idx, num := range nums {
        totalSum -= num
        if leftSum == totalSum {
            return idx
        }
        leftSum += num
    }
    return res
}