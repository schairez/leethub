//bottom up 2D tabular approach
//time: O(m*n)
//space: O(m*n)
func uniquePaths(m int, n int) int {
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
prev version
time: O(m*n) number of nodes in graph
space: O(n) for aux memo + O(n) for recursion stack
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