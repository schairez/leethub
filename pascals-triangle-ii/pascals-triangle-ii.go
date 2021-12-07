//time: O(n^2)
//space: O(n)
func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    var res []int
    for row := 0; row <= rowIndex; row++ {
        res = append(res, 1)
        for col := len(res)-2; col >=1; col-- {
            res[col] += res[col-1]
        }
    }
    return res
}
//version from 9/27/2021
//time: O(n^2)
// space: O(n^2)

func getRow2D(rowIndex int) []int {
    
    dp := make([][]int, rowIndex+1)
    for i:=0; i < len(dp); i++ {
        dp[i] = make([]int, i+1)
    }
    dp[0][0] = 1 
    if rowIndex == 0 {
        return dp[0]
    }
    dp[1][0], dp[1][1] = 1,1
    if rowIndex == 1 {
        return dp[1]
    }
    
    for i:=2; i<len(dp); i++ {
        dp[i][0] = 1 
        dp[i][len(dp[i])-1] = 1 
        for j :=1; j <len(dp[i])-1; j++ {
            dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
        }
    }
    
    return dp[rowIndex] 
    
}