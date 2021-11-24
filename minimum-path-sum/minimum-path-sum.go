
func min(a, b int) int { if a <= b { return a}; return b}



/*
DP 2D table approach

time: O(m*n)
space: O(m*n)
*/


func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    //case i == 0 && j == 0:
    //start
    dp[0][0] = grid[0][0]
    
    //from start row, we can only go right bound
    for col:=1; col < n; col++ {
        dp[0][col] = dp[0][col-1] + grid[0][col]
    }
    //from 0th col, start to bottom, can only go down bound
    for row:=1; row < m; row++ {
        dp[row][0] = dp[row-1][0] + grid[row][0]
    }
    for i := 1; i < m; i++ {
        for j:=1; j < n; j++ {
            dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
        }
    }
    return dp[m-1][n-1]
    
}

/*
TODO
DP 1D table approach

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j:=0; j < n; j++ {
        }
    }
}

*/