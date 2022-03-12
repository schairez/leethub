
//
//


func wordBreak(s string, wordDict []string) bool {
    n := len(s)
    dp := make([]bool, n+1)
    trie := &TrieNode{}
    for _, word := range wordDict {
        trie.InsertWord(word)
    }
    var currNode *TrieNode
    dp[0] = true
    for i := 0; i < n; i++ {
        if dp[i] {
            currNode = trie
            for j := i; j < n; j++ {
                if currNode == nil {
                    break
                }
                currNode = currNode.Children[s[j]- 'a']
                if currNode != nil && currNode.IsWord {
                    dp[j+1] = true
                }
            }
        } 
    } 
    return dp[n]
}

type TrieNode struct {
    Children [26]*TrieNode 
    IsWord   bool
}

func (t *TrieNode) InsertWord(word string) {
    for i := 0; i < len(word); i++ {
        if t.Children[word[i] - 'a'] == nil {
            t.Children[word[i] - 'a'] = &TrieNode{}  
        }
        t = t.Children[word[i] - 'a']
    }
    t.IsWord = true
}


/*

// dfs exhaustive search approach; TLE
// time: O(2^n)
// space: O(n)
func dfs(idx int, searchStr []byte, wordSet map[string]struct{}) bool {
    n := len(searchStr)
    if idx == n {
        return true
    }
    for i := idx+1; i <= n; i++ {
        if _, exists := wordSet[string(searchStr[idx:i])]; 
        exists && dfs(i, searchStr, wordSet) {
            return true
        }
    }
    return false
}

func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]struct{}, len(wordDict))
    for _, word := range wordDict {
        wordSet[word] = struct{}{}
    }
    searchWord := []byte(s)
    return dfs(0, searchWord, wordSet)
    
}
*/
