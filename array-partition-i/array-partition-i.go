//time: O(nlogn + n) ~ O(nlogn)
//space: O(1)

import "sort"

func arrayPairSum(nums []int) int {
    n := len(nums)
    sort.Ints(nums) //nlogn time
    totSum := 0
    for i:=1; i <n; i+=2 {
        totSum += min(nums[i], nums[i-1])
    }
    return totSum
    
}

func min(a,b int) int {
    if a <= b { return a}
    return b
}