/*
an arithmetic sequence can be added with 
Gaussian-arithmetic strategy n(n+1)/2
then we subtract the arithmetic sum and our sum of inputArr
the value remaining is the one missing
time: O(n)
space: O(1)
*/

func missingNumber(nums []int) int {
    n := len(nums)
    sequenceSum := n*(n+1) / 2 
    sumInputNums :=0
    for i :=0; i < n ; i ++ {
        sumInputNums += nums[i]
    }
    return sequenceSum - sumInputNums
}
