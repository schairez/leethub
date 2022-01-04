//BFS approach
//time: O(m*n)
//space: O(m*n)

func updateMatrix(mat [][]int) [][]int {
    var queue [][2]int  
    maxInt32 := (1 << 31) - 1
    numRows := len(mat)
    numCols := len(mat[0])
    distMatrix := make([][]int, numRows)
    for row := range distMatrix {
        distMatrix[row] = make([]int, numCols)
    }
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            if mat[r][c] == 1 {
                distMatrix[r][c] = maxInt32 
            } else {
                queue = append(queue, [2]int{r,c})
            }
        }
    }
    //N, E, S, W
    dirX := [4]int{1, 0, -1, 0}
    dirY := [4]int{0, 1,  0, -1}
    
    for len(queue) != 0 {
        pollNode := queue[0]
        queue = queue[1:]
        cellX, cellY := pollNode[0], pollNode[1]
        for i := range dirX {
            dx, dy  := dirX[i], dirY[i]
            cell2X, cell2Y := cellX + dx, cellY + dy
            if !(cell2X >= 0 && cell2X < numRows &&
                 cell2Y >= 0 && cell2Y < numCols) ||
                distMatrix[cell2X][cell2Y] <= distMatrix[cellX][cellY] + 1 {
                continue
            }
            queue = append(queue, [2]int{cell2X, cell2Y})
            distMatrix[cell2X][cell2Y] = 1 + distMatrix[cellX][cellY]
        }
    }
    return distMatrix
}