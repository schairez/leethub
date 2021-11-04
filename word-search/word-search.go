
// time: O(2^n)
// space: O(n)
/*

*/

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

func exist(board [][]byte, word string) bool {
    wordBytes := []byte(word)    
    visited := make(map[string]bool, len(word))
    //m*n grid
    m := len(board) //rows
    n := len(board[0]) //cols
    var res bool
    for i:=0; i < m; i++ {
        for j:=0; j < n; j++ {
            if board[i][j] == wordBytes[0] {
                if res { return true }
                dfs(&res, board, wordBytes, visited, 0, i, j, m, n)
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


func dfs(res *bool, board [][]byte, wordBytes []byte, visited map[string]bool, curByte, i, j, m, n int ) {
    rowColStr := string(i) + "_" + string(j)
    if !isValid(i, j, m, n) || visited[rowColStr] {
        return 
    }
    if wordBytes[curByte] != board[i][j] {
        return 
    }
    if curByte == len(wordBytes)-1 && wordBytes[curByte] == board[i][j] {
        *res = true
        return
    }
    visited[rowColStr] = true
    
    dirs := [4][2]int{
        {-1, 0}, //left
        {1, 0}, //right
        {0, 1}, //up
        {0, -1}, //down
    } 
    for _, dir := range dirs {
        curR, curC := i + dir[0], j + dir[1]
        if !isValid(curR, curC, m, n) { continue }
        dfs(res, board, wordBytes, visited, curByte+1, curR, curC, m, n)
    }
    visited[rowColStr] = false
    
}
    


