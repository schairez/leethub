/*
adjacent only 4 cardinal dirs

     1
     ^
 1 < B > 1
     V
     1
*/
// 1 <= m, n <= 50
// if click == mine; change to revealed min and return board
// elif click == emptyButAdjMine; figure out how many mines are adjacent and change to byte num rep
// else:
//   traverse neiNodes and change E to B recursively


    
func updateBoard(board [][]byte, click []int) [][]byte {
    const (
        unrevealedMine  = 'M'
        revealedMine    = 'X'
        unrevealedBlank = 'E'
        revealedBlank   = 'B'
    )
    r, c := click[0], click[1]
    if board[r][c] == unrevealedMine {
        board[r][c] = revealedMine
        return board
    }
    //mines := [8]byte{'1','2','3','4','5','6','7','8'}
    dirs := [8][2]int{
        [2]int{-1, 0},   //N
        [2]int{-1, 1},   //NE
        [2]int{0, 1},    //E
        [2]int{1, 1},   //SE
        [2]int{1, 0},    //S
        [2]int{1, -1},  //SW
        [2]int{0, -1},  //W
        [2]int{-1, -1}, //NW
    }
    numR, numC := len(board), len(board[0])
    inArea := func(r, c int) bool {
        return r >= 0 && r < numR && c >= 0 && c < numC
    }
    numAdjMines := func(r, c int) int {
        cnt := 0
        for _, dir := range dirs {
            nextR, nextC := r + dir[0], c + dir[1]
            if inArea(nextR, nextC) && board[nextR][nextC] == unrevealedMine {
                cnt++
            }
        }
        return cnt
    }
    //get num mines adj to click pt
    if cntMines := numAdjMines(r,c); cntMines != 0 {
        board[r][c] = byte(cntMines) + byte('0')
        return board
    }
    
    var (
        node          [2]int
        queue         [][2]int
        nextR, nextC  int
    )
    board[r][c] = revealedBlank
    queue = append(queue, [2]int{r, c})
    for len(queue) != 0 {
        node, queue = queue[0], queue[1:]
        r, c = node[0], node[1]
        for i := 0; i < 8; i++ {
            nextR, nextC = r + dirs[i][0], c + dirs[i][1]
            if !inArea(nextR, nextC) || 
            board[nextR][nextC] != unrevealedBlank {
                continue
            }
            
            //surrounding mines
            if cntMines := numAdjMines(nextR, nextC); cntMines != 0 {
                board[nextR][nextC] = byte(cntMines) + byte('0')
                continue
            }
            board[nextR][nextC] = revealedBlank
            queue = append(queue, [2]int{nextR, nextC})
        }
    }
    
    return board
    
}