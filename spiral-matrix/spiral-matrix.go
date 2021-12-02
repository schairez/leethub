//time: O(numR*numC)
//space: O(1)
//clockwise traversal 

func spiralOrder(matrix [][]int) []int {
    numR, numC := len(matrix), len(matrix[0])
    totalCells := numR * numC
    res := make([]int, totalCells)
    currCell := 0
    left, right := 0, numC-1
    top, bottom := 0, numR-1
    traverseCond := func() bool { return currCell < totalCells } 
    for traverseCond() {
        //traverse starting at top left to top right
        for col := left; col <= right && traverseCond(); col++ {
            res[currCell] = matrix[top][col]
            currCell++
        }
        top++
        //traverse from top right to bottom right
        for row := top; row <= bottom && traverseCond(); row++ {
            res[currCell] = matrix[row][right]
            currCell++
        }
        right--
        //traverse bottom right to bottom left
        for col := right; col >= left && traverseCond(); col-- {
            res[currCell] = matrix[bottom][col] 
            currCell++
        }
        bottom--
        //traverse bottom left to top left
        for row := bottom; row >= top && traverseCond(); row-- {
            res[currCell] = matrix[row][left]
            currCell++
        }
        left++
    }
    return res
} 