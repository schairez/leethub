//time: O(numR*numC); we iterate through all matrix cells
//space: O(1); no aux d.s. used

func matrixReshape(mat [][]int, r int, c int) [][]int {
    numR, numC := len(mat), len(mat[0])
    totalNums := numR*numC
    if totalNums != r*c {
        return mat
    }
    reshapedMatrix := make([][]int, r)
    for i := range reshapedMatrix {
        reshapedMatrix[i] = make([]int, c)
    }
    for numCnt:=0; numCnt < totalNums; numCnt++ {
        oldR, oldC := numCnt /numC, numCnt%numC
        newR, newC := numCnt/c, numCnt%c
        reshapedMatrix[newR][newC] = mat[oldR][oldC]
    } 
    return reshapedMatrix
} 
    
    