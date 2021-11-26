//prefix sum with hashmap approach
//constraint: 1 <= nums.length <= 2*10^4
//O(n) time and space
func subarraySum(nums []int, k int) int {
    //cnt continuous subarrays with sum = k
    cnt := 0
    runningSum := 0
    sumMap := make(map[int]int)
    sumMap[0] = 1 //cnt of sum 0 is 1
    for i:=0; i < len(nums); i++ {
        v := nums[i]
        runningSum += v
        if _, ok :=  sumMap[runningSum-k]; ok {
            cnt += sumMap[runningSum-k]
        }
        sumMap[runningSum]++
    }
    return cnt
}