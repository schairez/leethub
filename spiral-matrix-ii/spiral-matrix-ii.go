//time: O(n^2)
//space: O(n^2) since output grows as n grows
// 1 <= n <= 20
func generateMatrix(n int) [][]int {
    //totalCells := n*n
    cellCnt := 0
    matrix := make([][]int, n)
    for row := range matrix {
        matrix[row] = make([]int, n)
    }
    top, bottom := 0, n-1
    left, right := 0, n-1
    for top <= bottom && left <= right { 
        //traverse top left to top right
        for col := left; col <= right; col++ {
            matrix[top][col] = cellCnt+1
            cellCnt++
        }
        top++
        //traverse from top right to bottom right
        for row := top; row <= bottom; row++ {
            matrix[row][right] = cellCnt+1
            cellCnt++
        }
        right--
        //traverse from bottom right to bottom left
        for col := right; col >= left; col-- {
            matrix[bottom][col] = cellCnt+1
            cellCnt++
        }
        bottom--
        //traverse from bottom left to top left
        for row := bottom; row >= top; row-- {
            matrix[row][left] = cellCnt+1
            cellCnt++
        }
        left++
    }
    return matrix
}