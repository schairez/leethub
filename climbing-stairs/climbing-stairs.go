//constraint
//1 <= n <= 45

// three approaches
// 1. recursive memoized top-down
// 2. dp table bottom-up
// 3. two-pointer bottom-up

//recursive relation
// f(n) = f(n-1) + f(n-2)
// if n == 1 || n == 0 return 1

// 1. recursive memoized top-down
//space: O(n); time: O(n)
func climbStairsMemo(n int) int {
    return helper(n, nil)
}

func helper(n int, memo map[int]int) int {
    if n == 1 || n == 0 { 
        return 1
    }
    if memo == nil {
        memo = make(map[int]int)
    } else {
        if _, ok := memo[n]; ok {
            return memo[n]
        } 
    }
    memo[n] = helper(n-1, memo) + helper(n-2, memo)
    return memo[n]
}

// dp table bottom-up
//space: O(n); time: O(n)
func climbStairsDP(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for stair := 2; stair <= n; stair++ {
        dp[stair] = dp[stair-1] + dp[stair-2]
    }
    return dp[n]
}
//to derive f(n) we can keep track of three vars
// ptr1 = n-1; ptr2 = n-2; and the result var as the sum of the prev two
//we only need to keep track of the two previous values
//optimized bottom-up; maintaining two variables
func climbStairs(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    var (
        ptr1 = 1
        ptr2 = 1
    ) 
    res := ptr1 + ptr2
    for stair := 2; stair <= n; stair++ {
        res = ptr1 + ptr2
        ptr2, ptr1 = res, ptr2
    }
    return res
}








/*
//prev version from august

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i < n+1; i++ {
		third := first + second
		first = second
		second = third
	}
	return second

}
*/