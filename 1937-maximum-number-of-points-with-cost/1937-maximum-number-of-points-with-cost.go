/*
time: O(numR*numC)
space: O(numC)

*/

func max(a, b int) int { if a >= b { return a}; return b }

func maxPoints(points [][]int) int64 {
    var maxVal int64
    numRows := len(points)
    numCols := len(points[0])
    if numRows == 1 {
        for _, val := range points[0] {
            if int64(val) > maxVal {
                maxVal = int64(val)
            } 
        }
        return maxVal
    }
    //max coords at level
    dp := make([]int, numCols)
    for idx, pt := range points[0] {
        dp[idx] = int(pt)
    }
    fromLeft := make([]int, numCols)
    fromRight := make([]int, numCols)
    for row := 1; row < numRows; row++ {
        prev := make([]int, numCols)
        copy(prev, dp)
        fromLeft[0] = prev[0]
        for col := 1; col < len(points[row]); col++ {
            fromLeft[col] = max(fromLeft[col-1]-1,
                                prev[col],
                               )
        }
        fromRight[numCols-1] = 0
        for col := numCols-2; col >= 0; col-- {
            fromRight[col] = max(fromRight[col+1]-1,
                                 prev[col+1]-1)
        }
        for col := 0; col < numCols; col++ {
            dp[col] = max(fromLeft[col], 
                          fromRight[col]) + int(points[row][col])
        }
    }  
    //max number of pts with cost candidates
    for _, val := range dp {
        if int64(val) > maxVal {
            maxVal = int64(val)
        } 
    }
    return maxVal
}