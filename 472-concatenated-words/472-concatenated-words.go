
// 472. Concatenated Words
// trie + dfs approach
// w = maxLenWord
// n = numWords
// time: O(n) words to iter times O(w) to insertWord ≈ O(n*w)
// then to search for concat words O(n^2*w) 
// thus time ≈ O(n^2*w)
// space : O(w) to build trie

func findAllConcatenatedWordsInADict(words []string) []string {
    root := &TrieNode{}
    for _, word := range words {
        // 0 <= words[i].length <= 30
        if len(word) == 0 {
            continue
        }
        root.InsertWord(word)
    }
    var res []string
    
    for _, word := range words {
        if len(word) == 0 {
            continue
        }
        if isConcat(root, word, 0, 0) {
            res = append(res, word)
        }
    }
    return res
}

func isConcat(trie *TrieNode, word string, wordIdx, wordCnt int) bool {
    n := len(word)
    if wordIdx == n {
        return wordCnt >= 2
    }
    currNode := trie
    for i := wordIdx; i < n; i++ {
        idx := int(word[i] - 'a')
        if currNode.Children[idx] != nil {
            currNode = currNode.Children[idx]
            if currNode.IsWord && isConcat(trie, word, i+1, wordCnt+1) {
                return true
            }
        } else {
            return false
        }
    }
    return false
}

type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
}

func (trie *TrieNode) InsertWord(word string) {
    n := len(word)
    for i := 0; i < n; i++ {
        idx := int(word[i] - 'a')
        if trie.Children[idx] == nil {
            trie.Children[idx] = &TrieNode{}
        }
        trie = trie.Children[idx]
    } 
    trie.IsWord = true
}

func (trie *TrieNode) SearchWord(word string) bool {
    n := len(word)
    for i := 0; i < n; i++ {
        idx := int(word[i] - 'a')
        if trie.Children[idx] == nil {
            return false
        }
        trie = trie.Children[idx]
    }
    return trie.IsWord
}

