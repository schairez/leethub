/*
time: O(n)
space: O(1)
*/

import (
	"fmt"
)

func max(a, b int) int{
    if a < b {
        return b
    }
    return a 
}


func canJump(nums []int) bool {
    n := len(nums) 
    if n == 1 {
        return true 
    }
    furthest := 1 
    for i :=1; i < n; i++ {
        if furthest < i {
            return false 
        }
	v_end := nums[i-1] 
        if i +v_end >= n {
            //fmt.Println(dp)
            fmt.Printf("%d + %d >= n \n", i, v_end)
            return true 
        } else if v_end == 0 {
            continue 
        }
        furthest = max(furthest, v_end + i)
	
    }
    return furthest >= n
    //return dp[n]

}