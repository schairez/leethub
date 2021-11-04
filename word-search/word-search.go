
//naive backtracking approach
// time: O(N * 3^L); N is number of cells, L is length of word to search; 3 since we ignore the prev cell
// space: O(L); L is length of word to be matched

//TODO: faster approach with search pruning 


func exist(board [][]byte, word string) bool {
    wordBytes := []byte(word)    
    visited := make(map[string]bool, len(word))
    
    dirs := [4][2]int{
        {-1, 0}, //left
        {1, 0}, //right
        {0, 1}, //up
        {0, -1}, //down
    } 
    
    //m*n grid
    m := len(board) //rows
    n := len(board[0]) //cols
    
    var res bool
    for x:=0; x < m; x++ {
        for y:=0; y < n; y++ {
            if board[x][y] == wordBytes[0] {
                if res { return true }
                dfs(&res, dirs, board, wordBytes, visited, 0, x, y, m, n)
            }  
        }
    }
    
    return res
}


func isValid(curRow, curCol, numRows, numCols int) bool {
    if curRow < 0 || curRow > numRows -1 || curCol < 0 || curCol > numCols -1 {
        return false
    }
    return true 
} 


func dfs(res *bool, dirs [4][2]int, board [][]byte, wordBytes []byte, visited map[string]bool, curByte, x, y, m, n int ) {
    rowColStr := string(x) + "_" + string(y)
    if !isValid(x, y, m, n) || visited[rowColStr] {
        return 
    }
    if wordBytes[curByte] != board[x][y] {
        return 
    }
    if curByte == len(wordBytes)-1 && wordBytes[curByte] == board[x][y] {
        *res = true
        return
    }
    visited[rowColStr] = true
    
    for _, dir := range dirs {
        dx, dy := dir[0], dir[1]
        curR, curC := x + dx, y + dy
        if !isValid(curR, curC, m, n) { continue }
        dfs(res, dirs,board, wordBytes, visited, curByte+1, curR, curC, m, n)
    }
    visited[rowColStr] = false
    
}
    

/*


[["A","B","S","G"]
["S","F", "E","D"]
[["A","B","C","E"]

word = "ABCCED"
 ^
<->
 v
[["A","B","C","E"]
["S","F","C","S"]
["A","D","E","E"]]

@(0,0); 'A' == wordBytes[0] 
visited['A'] = true
                
                    /         \
                 dfs(-1, 0)    dfs(1,0)
       dfs**right
       dfs(1, i->1, 0) ; board[1][0] == wordBytes[1] 
       visited['B'] = true
       
       right again; i->2; j ->0;
       
    for i := range wordBytes {
        if !visited[wordBytes[i]] {
            return false
        }
    }
    return true

*/


