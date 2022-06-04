type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
}

func (trie *TrieNode) InsertWord(word string) {
    n := len(word)
    for i := 0; i < n; i++ {
        if trie.Children[word[i]-'a'] == nil {
            trie.Children[word[i]-'a'] = &TrieNode{}
        }
        trie = trie.Children[word[i]-'a']
    }
    trie.IsWord = true
}

func (trie *TrieNode) SearchWord(word string) bool {
    n := len(word)
    for i := 0; i < n; i++ {
        if trie.Children[word[i]-'a'] == nil {
            return false
        }
        trie = trie.Children[word[i]-'a']
    }
    return true
}


func wordBreak(s string, wordDict []string) []string {
    n := len(s)
    root := &TrieNode{}
    for _, word := range wordDict {
        root.InsertWord(word)
    }
    var (
        res []string
        dfs func(*TrieNode, []byte, int)
    )
    hashSet := make(map[string]struct{})
    dfs = func(node *TrieNode, currWord []byte, idx int) {
        if idx == n {
            if !node.IsWord {
                return
            }
            if currWord[len(currWord)-1] == ' ' {
                currWord = currWord[:len(currWord)-1]
            }
            cand := string(currWord)
            if _, exists := hashSet[cand]; !exists {
                res = append(res, cand)
                hashSet[cand] = struct{}{}
            }
            return
        }
        if childNode := node.Children[s[idx]-'a']; childNode != nil {
            currWord = append(currWord, s[idx])
            if childNode.IsWord {
                dfs(root, append(currWord, ' '), idx+1)
            } 
            dfs(childNode, currWord, idx+1)
            currWord = currWord[:len(currWord)-1]
        }
    }
    
    dfs(root, []byte{}, 0)
    return res
    
}