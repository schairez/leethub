

/*
Input: 
board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], 

words = ["oath","pea","eat","rain"]
        w := []byte(word)
        curr = root
        for _, wi := w {
            idx := wi - 'a'
            if curr.children[idx] == nil {
                curr.children[idx] = &Trie{}
            }
            curr = curr.children[idx]
        }
        curr.isWord = true

*/

type Trie struct {
    children [26]*Trie
    isWord bool
    wordsArrIdx int
    
}
func (t *Trie) Insert(word string, idx int) {
    wChar := []byte(word)
    curr := t
    for i:=0; i <len(wChar); i++ {
        idx := wChar[i] - 'a'
        if curr.children[idx] == nil { curr.children[idx] = &Trie{} }
        curr = curr.children[idx]
    }
    curr.isWord = true 
    curr.wordsArrIdx = idx
    return
}

func NewTrieFromWords(words []string) *Trie {
    root := &Trie{}
    for idx, word := range words {
        root.Insert(word,idx)
    }
    return root
} 
/*
visited []
n,s,e,w
  


*/
func dfs(words []string, trie *Trie, wordMap map[string]bool, board [][]byte, visited [][]bool, curRow, curCol, numRows, numCols int) {
    if curRow > numRows-1 || curRow < 0 || curCol < 0 || curCol > numCols-1  {
        return
    }
    if visited[curRow][curCol] == true { return }
    b := board[curRow][curCol]
    idx := b - 'a'
    var trieChild *Trie
    trieChild = trie.children[idx]
    if trieChild == nil { return }
    visited[curRow][curCol] = true
    
    if trieChild.isWord == true {
        //tmpBytes = append(tmpBytes, b)
        //tmpWord := make([]byte, len(tmpBytes))
        //copy(tmpWord, tmpBytes)
        wordMap[string(words[trieChild.wordsArrIdx])] = true
        
        //maxRes := cap(*res)
        //wordMap[string(tmpBytes)] = true
        //*res = append(*res, string(tmpBytes))
    }
    
    dirs := [4][2]int{
        {1, 0}, //mv east
        {-1, 0}, //mv west
        {0, 1}, //mv north
        {0, -1}, // mv south
    }
    //trie = trieChild
    var rowNext, colNext int 
    for _, dir := range dirs {
        rowNext, colNext = curRow + dir[0], curCol + dir[1]
        dfs(words, trieChild, wordMap, board, visited, rowNext, colNext,numRows,numCols)
        //dfs(trie, res, append(tmpBytes,b), board, visited, rowNext, colNext,numRows,numCols, cap(*res))
    }
    visited[curRow][curCol] = false
    
}


func findWords(board [][]byte, words []string) []string {
    numRows := len(board)
    numCols := len(board[0])
    visited := make([][]bool, numRows)
    for i := range visited {
        visited[i] = make([]bool, numCols)
    }
    
    trie := NewTrieFromWords(words)
    set := make(map[string]bool, len(words))
    res := make([]string, 0, len(words))
    var b byte
    for i:=0; i < numRows; i++ {
        for j:=0; j < numCols; j++ {
            b = board[i][j] 
            idx := b - 'a'
            if trie.children[idx] != nil {
                dfs(words, trie, set, board, visited, i, j, numRows, numCols)
            }
        }
    }
    for validWord, _ := range set {
        res = append(res, validWord)
    }
    
    return res
}





