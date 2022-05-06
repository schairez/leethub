
func longestWord(words []string) string {
    root := &TrieNode{}
    for _, word := range words {
        root.InsertWord(word)
    }
    var (
        res string
        chars []byte
        dfs func(*TrieNode, int)
    )
    dfs = func(trie *TrieNode, currLen int) {
        if currLen > len(res) {
            res = string(chars)
        } else if currLen == len(res) {
            cand := string(chars)
            if cand < res {
                res = cand
            }
        }
        for idx, childNode := range trie.Children {
            if childNode != nil && childNode.IsWord {
                chars = append(chars, byte(idx)+'a')
                dfs(childNode, currLen+1)
                chars = chars[:len(chars)-1]
            }
        }
    }
    
    dfs(root, 0)
    
    return res
}


type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
}

func (trie *TrieNode) InsertWord(word string) {
    n := len(word)
    for i := 0; i < n; i++ {
        idx := int(word[i]-'a')
        if trie.Children[idx] == nil {
            trie.Children[idx] = &TrieNode{}
        }
        trie = trie.Children[idx]
    }
    trie.IsWord = true
}

