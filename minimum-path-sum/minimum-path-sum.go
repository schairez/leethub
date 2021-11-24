/*

*/

func min(a, b int) int { if a <= b { return a}; return b}

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j:=0; j < n; j++ {
            switch {
            //start
            case i == 0 && j == 0:
                dp[i][j] = grid[i][j]
            //from start row, we can only go right bound
            case i == 0 && j >= 1:
                dp[i][j] = dp[i][j-1] + grid[i][j]
            //from 0th col, start to bottom, can only go down bound
            case i > 0 && j == 0: 
                dp[i][j] = dp[i-1][j] + grid[i][j]
            default:
                dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
            }
        }
    }
    //fmt.Println(dp)
    return dp[m-1][n-1]
    
}
