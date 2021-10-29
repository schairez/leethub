// m -> partTarget ~ n/2
//space: O(m) ~ O(m)
//time: O(n*m)

func canPartition(nums []int) bool {
    n := len(nums)
    total := 0
    for _, v := range nums {
        total += v
    }
    if total % 2 != 0 {
        return false
    }
    partTarget := total / 2
    dp := make([]bool, partTarget+1 )
    dp[0] = true
    for i:=0; i < n; i++ {
        for j:=partTarget; j >= nums[i]; j-- {
            dp[j] = dp[j] || dp[j-nums[i]]
        }
    }
    return dp[partTarget]
}
/*
nums = [1,5,11,5]
if sum % 2 != 0 return false
partTarget = sum(nums)/2 -> 11
dp [0..11] in {T,F}
dp[0] -> T; choose none

*/
