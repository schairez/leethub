
/*
In chess, a queen can attack horizontally, vertically, and diagonally
*/
// 1 "Q" per row
// moment we have no available spaces we backtrack

func solveNQueens(n int) [][]string {
    var res [][]string 
    if n == 1 {
        res = append(res, []string{"Q"})
        return res
    } 
    board := buildBaseBoard(n)
    solve(&res, board, 0, n, n)
    return res
}
func buildBaseBoard(n int) [][]byte {
    delim := byte('.')
    base := make([]byte, n)
    for i:=0; i < len(base); i++ {
        base[i] = delim
    }
    baseBoard := make([][]byte, n)
    for i := range baseBoard {
        baseBoard[i] = make([]byte, n) 
        copy(baseBoard[i], base)
    }
    
    return baseBoard
}


//checks if valid loc at desired row,col val
func isValid(board [][]byte, row, col, sideLen int) bool {
    //check above curr row
    for i:=0; i < row; i++ {
        if board[i][col] == byte('Q') {
            return false
        }
    }
    //check upper right diagonal
    for i, j := row-1, col+1; i >=0 && j < sideLen; i, j = i -1, j+1 {
        if board[i][j] == byte('Q') {
            return false
        } 
    }
    //check upper left diagonal
    for i, j := row-1, col-1; i >=0 && j >=0; i, j = i-1, j-1 {
        if board[i][j] == byte('Q') {
            return false
        }
    }
    
    return true
    
}




func solve(res *[][]string, board [][]byte, currRow int, nLeft int, sideLen int) {
    if nLeft == 0 {
        tmp := make([]string, sideLen)
        //idx :dd= 0
        for i:=0; i <len(board); i++ {
            //copy(tmp[i], string(board[i]))
            tmp[i] = string(board[i])
        }
        *res = append(*res, tmp)
        return 
    }
    for col:=0; col < sideLen; col++ {
        if !isValid(board, currRow, col, sideLen) {
            continue
        }
        board[currRow][col] = byte('Q')
        solve(res, board,currRow+1, nLeft-1, sideLen)
        board[currRow][col] = byte('.')
    }
} 

/*

if n == 2; []string{"Q.", ".Q"}
if n == 3; []string{"Q..", ".Q.", "..Q"}
if n == 4; []string{"Q...", ".Q..", "..Q.", "...Q"}


func buildOpts(n int) []string {
    queen := "Q"
    delim := "."
    var opts []string
    str := strings.Repeat(delim, n)
    opts = append(opts, queen + str[:len(str)-1])
    for i := 1; i < n-1; i++ {
        opt := str[:i] + queen + str[i+1:] 
        opts = append(opts, opt)
    }
    lastOpt := strings.Repeat(delim, n-1) + queen
    opts = append(opts, lastOpt)
    return opts
}
*/

