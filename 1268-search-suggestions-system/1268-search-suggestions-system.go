
type TrieNode struct {
    Children [26]*TrieNode
    Char byte
    IsWord bool
}

func (t *TrieNode) InsertWord(word string) {
    for i := 0; i < len(word); i++ {
        idx := int(word[i]-'a')
        if t.Children[idx] == nil {
            t.Children[idx] = &TrieNode{}
            t.Children[idx].Char = word[i]
        }
        t = t.Children[idx]
    }
    t.IsWord = true
}

func dfs(res *[]string, seq []byte, trie *TrieNode) {
    n := len(*res)
    if n == 3 {
        return
    } 
    if trie == nil {
        return
    }
    if trie.IsWord {
        *res = append(*res, string(seq))
    }
    for i := 0; i < 26; i++ {
        childNode := trie.Children[i]
        if childNode != nil {
            dfs(res, append(seq, childNode.Char), childNode)   
        }
    }
}

func findWords(trie *TrieNode, prefix string) []string {
    if trie == nil {
        return []string{}
    }
    for i := 0; i < len(prefix); i++ {
        if trie.Children[prefix[i]-'a'] == nil {
            return []string{}
        }
        trie = trie.Children[prefix[i]-'a']
    }
    res := make([]string, 0, 3)
    dfs(&res, []byte(prefix), trie)
    
    return res
}

func suggestedProducts(products []string, searchWord string) [][]string {
    trie := &TrieNode{}
    for _, prod := range products {
        trie.InsertWord(prod)
    }
    res := make([][]string, 0, len(products))
    n := len(searchWord)
    for i := 0; i < n; i++ {
        res = append(res, findWords(trie, searchWord[:i+1]))
    }
                 
    return res
}

