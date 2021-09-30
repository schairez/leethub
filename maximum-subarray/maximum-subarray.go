/*
time: O(n)
space: O(1)
*/

func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    //local and global Sums
    gSum, lSum := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        currVal := nums[i]
        lSum = max(lSum + currVal, currVal)
        gSum = max(gSum, lSum)
    }
    return gSum
}

func max(a,b int) int {
    if a > b {
        return a 
    }
    return b 
}