func min(a, b int) int { if a <= b { return a}; return b}

//recurrence formula
//dp[r][c] = grid[r][c] + min{from up, from left}
/*
DP 2D table approach

time: O(m*n)
space: O(m*n)
*/

func minPathSum2Ddp(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    dp := make([][]int, numR)
    for row := range dp {
        dp[row] = make([]int, numC)
    }
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
            switch {
            case r == 0 && c == 0: //start
                dp[r][c] = grid[r][c]
            case r == 0: //from left only
                dp[r][c] = dp[r][c-1] + grid[r][c]
            case c == 0: //from above only
                dp[r][c] = dp[r-1][c] + grid[r][c] 
            default:
                fromAbove := dp[r-1][c]
                fromLeft := dp[r][c-1]
                dp[r][c] = min(fromAbove, fromLeft) + grid[r][c]
            }
        }
    }
    return dp[numR-1][numC-1]
}

//DP 1D table approach
//time: O(m*n)
//space: O(n)
//solution is at origin
//we move the opposite dir

func minPathSum(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    dp := make([]int, numC)
    for r := numR-1; r >= 0; r-- {
        for c := numC-1; c >= 0; c-- {
            switch {
            //from right to left only
            //if base-row and not base-col    
            case r == numR-1 && c != numC-1:
                dp[c] = dp[c+1] + grid[r][c] 
            //moving upward
            //if base-col, but from down to up row
            case r != numR-1 && c == numC-1:
                dp[c] = grid[r][c] + dp[c]
            //move from right to left
            //then, we repeat as we move from upwards
            //from below row
            //if non-base row and non-base col
            case c != numC-1 && r != numR-1:
                dp[c] = grid[r][c] + min(dp[c], dp[c+1])
                
            default:
            //start node case
            //case r == numR-1 && c == numC-1:
                dp[c] = grid[r][c]
            }
        }
    }
    return dp[0]
}
        








/*
//version from 11/23/2021
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
*/

