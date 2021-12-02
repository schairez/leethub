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
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

/*
func spiralOrder(matrix [][]int) []int {
    numR, numC := len(matrix), len(matrix[0])
    res := make([]int, numR*numC)
    cell := 0
    curR, curC := 0, 0
    topR, bottomR, leftC, rightC := 0, numR-1, 0, numC -1
    for cell < numR * numC {
        //go right
        iterC := curC
        for ;iterC <= rightC && cell < numR *numC; iterC++ {
            res[cell] = matrix[curR][iterC]
            cell++
        }
        topR++
        
        curC = iterC
        //go down
        iterR := curR
        for ; iterR <= bottomR && cell < numR *numC; iterR++ {
            res[cell] = matrix[iterR][curC]
            cell++
        }
        rightC--
        
        iterC = curC
        //go left; row fixed, col--
        for ; iterC >= leftC && cell < numR *numC; iterC-- {
            res[cell] = matrix[curR][iterC]
            cell++
        }
        curC = iterC
        bottomR--
        //go up
        for; iterR >= topR && cell < numR * numC; iterR-- {
            res[cell] = matrix[iterR][curC]
            cell++
        }
        leftC++
    }
    return res
    
}
*/