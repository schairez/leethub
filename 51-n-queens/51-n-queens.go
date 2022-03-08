
func solveNQueens(n int) [][]string {
    if n == 1 {
        return [][]string{{"Q"}}
    }
    if n == 2 || n == 3 {
        return [][]string{}
    }
    var res [][]string
    mat := genMatrix(n)
    dfs(mat, &res, 0, n)
    
    return res
}

func dfs(matrix [][]byte, res *[][]string, x int, queensLeft int) {
    if queensLeft == 0 {
        n := len(matrix)
        tmp := make([]string, n)
        for x := range matrix {
            tmp[x] = string(matrix[x])
        }
        *res = append(*res, tmp)
        return
    }
    for y := 0; y < len(matrix); y++ {
        if !isValid(matrix, x, y) {
            continue
        }
        matrix[x][y] = 'Q'
        dfs(matrix, res, x+1, queensLeft-1)
        matrix[x][y] = '.'
    }
}



func genMatrix(n int) [][]byte {
    rowVals := make([]byte, n)
    for idx := range rowVals {
        rowVals[idx] = '.'
    }
    matrix := make([][]byte, n)
    for i := 0; i < n; i ++ {
        //matrix[i] = rowVals
        matrix[i] = make([]byte, n)
        copy(matrix[i], rowVals)
    }
    return matrix
}

func isValid(matrix [][]byte, queenLocX, queenLocY int) bool {
    n := len(matrix)
    // check col wise
    for y := 0; y < n; y++ {
        if matrix[queenLocX][y] == 'Q' {
            return false
        }
    }
    // check row wise
    for x := 0; x < n; x++ {
        if matrix[x][queenLocY] == 'Q' {
            return false
        }
    }
    // offset +1 each time
    // both up and down dirs

    // diagonals down
    offset := 1
    // [0,0]; 
    for x := queenLocX+1; x < n; x++ {
        diag1 := queenLocY + offset
        if diag1 >= 0 && diag1 < n {
            if matrix[x][diag1] == 'Q' {
                return false
            }
        }
        diag2 := queenLocY - offset
        if diag2 >= 0 && diag2 < n {
            if matrix[x][diag2] == 'Q' {
                return false
            }
        }
        offset++
    }
    // diagonals up
    offset = 1
    for x := queenLocX-1; x >= 0; x-- {
        diag1 := queenLocY + offset
        if diag1 >= 0 && diag1 < n {
            if matrix[x][diag1] == 'Q' {
                return false
            }
        }
        diag2 := queenLocY - offset
        if diag2 >= 0 && diag2 < n {
            if matrix[x][diag2] == 'Q' {
                return false
            }
        }
        offset++
    }
    return true
}
    

