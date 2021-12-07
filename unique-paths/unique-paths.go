//recursive top-down version
//space: O(m*n)
//time: O(m*n)
func uniquePathsMemo(m int, n int) int {
    memo := make([][]int, m) 
    for row := range memo {
        memo[row] = make([]int, n)
    }
    var dfs func(int, int) int
    dfs = func(r, c int) int {
        if numPaths := memo[r][c]; numPaths !=0 {
            return numPaths
        } 
        switch {
        case r == 0 && c == 0:
            return 0
        case r == 0 || c == 0:
            return 1
        }
        memo[r][c] = dfs(r-1, c) + dfs(r, c-1)
        return memo[r][c]
    }
    return dfs(m-1, n-1)
}


//bottom up 2D tabular approach
//time: O(m*n)
//space: O(m*n)
func uniquePaths2D(m int, n int) int {
    dp := make([][]int, m)
    
    for i:=0; i < m; i++ {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    for j:=0; j < n; j++ {
        dp[0][j] = 1
    }
    for r:=1; r < len(dp); r++ {
        for c:=1; c < len(dp[0]); c++ {
            dp[r][c] = dp[r][c-1] + dp[r-1][c]
        }
    } 
    
    return dp[m-1][n-1]
}

/*
reduced to O(n) space bottom up DP version
- we only need to reserve space for each col in our dp table
- taking care of initializing our dp arr with vals = 1,
  this marks the iteration of traversing right bound on row0
- when iterating through each remaining (r,c) pair row-wise,
  we increment the store at that dp idx by the unique paths from the left
  of the dp idx; recurrence relation: dp[i] = dp[i] + dp[i-1] i elem of {1,n-1}
- after each row traversal, our dp array stores results for unique paths
  up to that row
for m=3, n=7 ex
[1, 1, 1, 1, 1, 1, 1] after row=0 traversal
[1, 2, 3, 4, 5, 6, 7] after row=1 traversal
[1, 3, 6, 10, 15, 21, 28] after row=2 traversal; res =28

time: O(m*n)
space: O(n)
*/

func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for idx := range dp {
        dp[idx] = 1
    }
    for row := 1; row < m; row++ {
        for col := 1; col < n; col++ {
            dp[col] += dp[col-1]
        } 
    }
    fmt.Println(dp)
    return dp[n-1]
}


/*
prev version; 11/22/2021
time: O(m*n) number of nodes in graph
space: O(m*n) for aux memo + O(n) for recursion stack
*/

/*

import (
    "strconv"
    "fmt"
)

func uniquePaths(m int, n int) int {
    memo := make(map[string]int)
    return dfs(memo, 0, 0, m, n )
}


func isValid(i, j, m, n int) bool {
    if i == m || j == n {
        return false
    }
    return true 
}

func dfs(memo map[string]int, i, j, m, n int) int {
    if !isValid(i,j,m,n) {
        return 0
    }
    key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
    if _, ok := memo[key]; ok {
        return memo[key]
    }
    if i == m -1 && j == n - 1 {
        return 1 
    }
    path := 0 
    path += dfs(memo,i+1, j, m, n)
    path  += dfs(memo,i, j+1, m, n)
    memo[key] = path
    fmt.Println(memo)
    return memo[key]
    
}
*/