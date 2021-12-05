//kadene's algo
//time: O(n); space: O(1)

//1 <= nums.length <= 10^5
//-10^4 <= nums[i] <= 10^4

func max(a, b int) int { if a >= b { return a }; return b}

func maxSubArray(nums []int) int {
    globalMaxSubArr := -1 << 31
    localMaxSubArr := globalMaxSubArr
    
    for _, num := range nums {
        localMaxSubArr = max(localMaxSubArr + num , num )
        globalMaxSubArr = max(globalMaxSubArr, localMaxSubArr)
    }
    return globalMaxSubArr
    
}



/*
2021-09-29 version
time: O(n)
space: O(1)

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
*/
