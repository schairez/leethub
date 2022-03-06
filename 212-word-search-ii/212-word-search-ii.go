// 
// 
//

type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
    WordIdx int
}

func NewTrie() *TrieNode {
    return &TrieNode{}
}

// contains word
func (trie *TrieNode) AddWord(word string, wordIdx int) {
    for i := 0; i < len(word); i++ {
        key := word[i] - 'a'
        if trie.Children[key] == nil {
            trie.Children[key] = &TrieNode{}
        }
        trie = trie.Children[key]
    }
    trie.IsWord = true
    trie.WordIdx = wordIdx
}





func findWords(board [][]byte, words []string) []string {
    numX, numY := len(board), len(board[0])
    dirs := [5]int{1,0,-1,0,1}
    wordMap := make(map[string]struct{})
    visited := make([][]bool, numX)
    for x := range visited {
        visited[x] = make([]bool, numY)
    }
    
    trieRoot := NewTrie()
    for idx, word := range words {
        trieRoot.AddWord(word, idx)
    }
    var dfs func(int, int, *TrieNode)
    dfs = func(x, y int, trie *TrieNode) {
        if x < 0 || x >= numX || y < 0 || y >= numY {
            return
        }
        if visited[x][y] {
            return
        }
        idx := board[x][y] - 'a'
        trie = trie.Children[idx]
        if trie == nil {
            return
        }
        visited[x][y] = true
        if trie.IsWord {
            wordMap[words[trie.WordIdx]] = struct{}{}
        }
        for i:=1; i < len(dirs); i++ {
            deltaX, deltaY := dirs[i], dirs[i-1]
            newX, newY := deltaX +x, deltaY + y
            dfs(newX, newY, trie)
        }
        visited[x][y] = false
    }
    
    
    for x := 0; x < numX; x++ {
        for y := 0; y < numY; y++ {
            dfs(x, y, trieRoot)
        }
    }
    res := make([]string, 0, len(wordMap))
    for wordKey := range wordMap {
        res = append(res, wordKey)
    }
    return res
}